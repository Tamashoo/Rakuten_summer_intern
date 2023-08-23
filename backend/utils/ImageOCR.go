package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
)

type EntityAnnotation = visionpb.EntityAnnotation

func ocr(filepath string) []*EntityAnnotation {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 5)
	if err != nil {
		log.Fatalf("Failed to detect texts: %v", err)
	}

	// 最初は全体をテキストにした結果が入るので飛ばす
	annotations = annotations[1:]
	return annotations
}

func getTotalPayment(annotations []*EntityAnnotation, target string, threshold float64) (int, error) {
	yInterval := [2]int32{-1, -1}
	var xBoundary int32 = -1
	for _, annotation := range annotations {
		if strings.Contains(annotation.Description, target) {
			yInterval[0] = annotation.BoundingPoly.Vertices[0].Y
			yInterval[1] = annotation.BoundingPoly.Vertices[3].Y
			xBoundary = annotation.BoundingPoly.Vertices[1].X
		}
	}
	paymentText := ""
	for _, annotations := range annotations {
		targetYInterval := [2]int32{annotations.BoundingPoly.Vertices[0].Y, annotations.BoundingPoly.Vertices[3].Y}
		targetX := annotations.BoundingPoly.Vertices[1].X
		if math.Abs((float64(targetYInterval[0]-yInterval[0]))) < threshold && math.Abs(float64(targetYInterval[1]-yInterval[1])) < threshold && targetX > xBoundary {
			paymentText += annotations.Description
		}
	}
	re := regexp.MustCompile(`[¥,.]`)
	paymentText = re.ReplaceAllString(paymentText, "")
	return strconv.Atoi(paymentText)
}

func fileToTexts(filepath string) []string {
	fp, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	texts := []string{}
	for scanner.Scan() {
		texts = append(texts, scanner.Text())
	}
	return texts
}

// 各商品の値引額がよくわからない場合
// 経験値として単純な値引額の合計を返す
func annotationsToPoint(annotations []*EntityAnnotation, threshold float64) int {
	discountSum := 0

	for _, annotation := range annotations {
		if strings.HasPrefix(annotation.Description, "-") {
			num, err := strconv.Atoi(annotation.Description[1:])
			if err == nil {
				discountSum += num
			}
		}
	}
	// total, err := getTotalPayment(annotations, "合計", threshold)
	// fmt.Println(total, discountSum)
	// if err != nil {
	// 	log.Fatalf("Failed to get total payment: %v", err)
	// }

	return discountSum * 100
}

// 各商品の値引き額がわかる場合
// 割引率の合計(経験値)と割引商品の数（貢献度）を返す
func calculatePoint(annotations []*EntityAnnotation, threshold float64) (int, int) {
	pts := 0
	cnt := 0
	// 左上のポイントのY座標で昇順ソート
	sort.Slice(annotations, func(i, j int) bool {
		return annotations[i].BoundingPoly.Vertices[0].Y > annotations[j].BoundingPoly.Vertices[0].Y
	})
	for idx, annotation := range annotations {
		if strings.HasPrefix(annotation.Description, "-") {
			discountValue, err := strconv.Atoi(annotation.Description[1:])
			xPosition := annotation.BoundingPoly.Vertices[0].X
			cnt++
			// -(数字) というものが見つかったら
			// その左上のポイントのY座標より小さいY座標を持つ数字の中で最も大きいY座標を持つものを見つける
			// すなわちそれより後に出てくる数字のうち最初に出てくるものをとる
			if err == nil {
				for _, underAnnotation := range annotations[idx:] {
					underXPosition := underAnnotation.BoundingPoly.Vertices[0].X
					listPrice, err := strconv.Atoi(underAnnotation.Description)
					if err == nil && listPrice > discountValue && math.Abs(float64(underXPosition-xPosition)) < threshold {
						fmt.Println(listPrice, discountValue)
						discountRate := 100 * discountValue / listPrice
						pts += discountRate
						break
					}
				}
			}
		}
	}
	return pts, cnt * 200
}

// エンコード
func encode(filename string) string {

	file, _ := os.Open(filename)
	defer file.Close()

	fi, _ := file.Stat() //FileInfo interface
	size := fi.Size()    //ファイルサイズ

	data := make([]byte, size)
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data)
}

// デコード
func decode(str string) string {
	data, _ := base64.StdEncoding.DecodeString(str) //[]byte

	file, _ := os.Create("encode_and_decord.png")

	decodepath := "encode_and_decord.png"
	defer file.Close()

	file.Write(data)

	return decodepath
}

func main() {
	enc_data := encode("スクリーンショット 2023-08-22 113438.png")
	decodepath := decode(enc_data)
	// fmt.Println(decodepath)
	annotations := ocr(decodepath)
	fmt.Println(calculatePoint(annotations, 10))
}