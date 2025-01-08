package parse

import (
	"bufio"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"log"
	"strconv"
	"strings"
	"time"
)

func WgDump(output string) ([]models.Peer, error) {
	var result []models.Peer
	scanner := bufio.NewScanner(strings.NewReader(output))

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t") // Split line by tab delimiter

		// Ensure fields count matches the expected number
		if len(fields) < 8 {
			log.Printf("invalid line format: %s", line)
			continue
		}

		dump := models.Peer{
			Interface:     fields[0],
			PublicKey:     fields[1],
			Port:          fields[2],
			Endpoint:      fields[3], // Update field index for Endpoint based on new format
			AllowedIPs:    fields[4], // Incorporating AllowedIPs from the new dump format
			LastHandshake: fields[5],
			TransferRx:    fields[6],
			TransferTx:    fields[7],
			Flags:         fields[8], // Incorporating Flags from the new dump format
		}

		// Handle optional AllowedIPs and Flags
		if len(fields) > 8 {
			dump.AllowedIPs = fields[8]
		}
		if len(fields) > 9 {
			dump.Flags = fields[9]
		}

		if dump.LastHandshake != "" && dump.LastHandshake != "0" {
			timestamp, err := strconv.ParseInt(dump.LastHandshake, 10, 64)
			if err == nil {
				dump.LastHandshake = time.Unix(timestamp, 0).Format("2006-01-02T15:04:05Z07:00")
			} else {
				log.Printf("error parsing LastHandshake as int: %v", err)
			}
		} else {
			dump.LastHandshake = ""
		}
		// Append the parsed struct to the list
		result = append(result, dump)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
