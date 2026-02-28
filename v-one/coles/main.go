package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// Create a new browser context with additional options
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("disable-javascript", false),
		chromedp.Flag("disable-web-security", false),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("user-data-dir", "./chrome-data"),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-plugins", true),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("disable-notifications", true),
		chromedp.Flag("enable-cookies", true),
		chromedp.Flag("enable-javascript", true),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// ctx, cancel := chromedp.NewContext(context.Background())
	// defer cancel()

	// Optional: add a timeout
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var pageContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.coles.com.au/product/daley-street-mediumdark-beans-1kg-5556286"),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.Sleep(time.Duration(2+rand.Intn(3))*time.Second), // Random delay between 2-4 seconds
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Simulate random mouse movements
			time.Sleep(time.Duration(1+rand.Intn(2)) * time.Second)
			return nil
		}),
		chromedp.WaitReady("document", chromedp.ByQuery),
		chromedp.WaitReady("window", chromedp.ByQuery),
		chromedp.OuterHTML("html", &pageContent),
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Fetched page content:")
	fmt.Println(pageContent)
}
