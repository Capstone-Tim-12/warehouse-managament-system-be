package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/http/core"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/sashabaranov/go-openai"
)

func (s *defaultUser) ChatBot(ctx context.Context, userId int, text string) (resp model.ChatResponse, err error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Halo! Perkenalkan aku adalah WMS Automated Assistant. Bagimana aku bisa membantumu hari ini?",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Bagaimana aku dapat melihat informasi gudang?",
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Untuk dapat melihat informasi gudang, anda dapat membuka beranda aplikasi lalu melihat list gudang yang tersedia dan klik gudang yang menurut kamu menarik, maka kamu akan melihat informasi detai gudangnya, informasi gudang meliputi nama gudang, lokasi gudang, harga sewa dan ukuran gudang",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Bisakah aku mengganti akun?",
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Mohon maaf, untuk menjamin keamanan akun kamu, kamu tidak bisa mengganti akun yang telah di daftarkan",
			},
			{
				Role: openai.ChatMessageRoleUser,
				Content: "Aku punya pertanyaan lain",
			},
			{
				Role: openai.ChatMessageRoleSystem,
				Content: "Silakan masukan pertanyaan kamu, aku akan selalu siap untuk menjawab pertanyaanmu",
			},
			{
				Role: openai.ChatMessageRoleUser,
				Content: "bagaimana cara melakukan sewa gudang?",
			},
			{
				Role: openai.ChatMessageRoleSystem,
				Content: "anda bisa melakukan order pada bagian detail gudang dan masukan cara cicilan dan durasi sewa, cicilan tersedia dalam jangka waktu mingguan, bulanan atau tahunan",
			},
			{
				Role: openai.ChatMessageRoleUser,
				Content: "jelaskan tentang aplikasi ini",
			},
			{
				Role: openai.ChatMessageRoleSystem,
				Content: "nama aplikasi ini adalah Digihouse, adalah aplikasi untuk melakukan sewa gudang bagi anda yang membutuhkan gudang untuk keperluan penjualan atau penyimpanan barang",
			},
		},
	}

	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: text,
	})

	key := fmt.Sprintf("ChatBot_%v", userId)
	ultiData, _ := s.coreRepo.GetUtilityData(ctx, key)
	if ultiData.Data.Value != "" {
		var reqData openai.ChatCompletionRequest
		err = json.Unmarshal([]byte(ultiData.Data.Value), &reqData)
		if err != nil {
			fmt.Println("error unmarshalling chat: ", err.Error())
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		reqData.Messages = append(reqData.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		req = reqData
	}

	dataResp, err := s.chatWrapper.GenerateText(ctx, req)
	if err != nil {
		fmt.Println("error generating text: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	req.Messages = append(req.Messages, dataResp.Choices[0].Message)
	dataByte, _ := json.Marshal(req)
	reqUlti := core.SetUtilityRequest{
		Key:      key,
		Value:    string(dataByte),
		Duration: 3600,
	}
	_, err = s.coreRepo.SetUtility(ctx, reqUlti)
	if err != nil {
		fmt.Println("error set utility: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	resp.Text = dataResp.Choices[0].Message.Content
	return
}
