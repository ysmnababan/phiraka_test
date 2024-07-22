package handler

import (
	"bytes"
	"encoding/base64"
	"image/color"
	"image/png"
	"math/rand"
	"net/http"
	"time"

	"github.com/fogleman/gg"
	"github.com/labstack/echo/v4"
)

var captchaStore = make(map[string]string)

// Generate a random CAPTCHA string
func generateCaptchaString(length int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	captcha := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range captcha {
		captcha[i] = letters[rand.Intn(len(letters))]
	}
	return string(captcha)
}

// Generate CAPTCHA image with noise and distortion
func generateCaptcha() (string, string) {
	const width, height = 200, 80
	const captchaLength = 6
	captchaStr := generateCaptchaString(captchaLength)

	dc := gg.NewContext(width, height)
	dc.SetColor(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	dc.Clear()

	// Add random noise
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		r := uint8(rand.Intn(256)) // Convert int to uint8
		g := uint8(rand.Intn(256)) // Convert int to uint8
		b := uint8(rand.Intn(256)) // Convert int to uint8
		dc.SetColor(color.RGBA{R: r, G: g, B: b, A: 255})
		dc.DrawCircle(rand.Float64()*width, rand.Float64()*height, 2)
		dc.Fill()
	}

	// Draw the text
	dc.SetColor(color.Black)
	fontSize := 48.0
	dc.LoadFontFace("/path/to/your/font.ttf", fontSize)
	dc.DrawStringAnchored(captchaStr, width/2, height/2, 0.5, 0.5)

	img := dc.Image()
	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	captchaBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	captchaID := base64.StdEncoding.EncodeToString(randBytes(16))
	captchaStore[captchaID] = captchaStr

	return captchaID, "data:image/png;base64," + captchaBase64
}

// randBytes generates a slice of random bytes of length n
func randBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}
	return b
}

// Handle CAPTCHA request
func GetCaptcha(c echo.Context) error {
	captchaID, captchaBase64 := generateCaptcha()
	return c.JSON(http.StatusOK, map[string]string{
		"captcha_id":    captchaID,
		"captcha_image": captchaBase64,
	})
}
