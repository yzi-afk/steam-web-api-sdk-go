# Steam Web API SDK for Go

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Status](https://img.shields.io/badge/status-work%20in%20progress-red)](Status)

## Example

   ```go
   import "github.com/yaihub/steam-web-api-sdk-go"

   func main() {
      httpClient := &http.Client{}
      client, err := steam.New("<YOUR_API_KEY>", httpClient)
      if err != nil {
        // Handle error
      }
      _ = client
    }
   ```