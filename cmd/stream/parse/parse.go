package parse

import (
	"bufio"
	"github.com/softwareplace/wireguard-api/cmd/stream/spec"
	"log"
	"strings"
)

func WgDump(output string) ([]spec.WgDump, error) {
	var result []spec.WgDump
	scanner := bufio.NewScanner(strings.NewReader(output))

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t") // Split line by tab delimiter

		// Ensure fields count matches the expected number
		if len(fields) < 8 {
			log.Printf("invalid line format: %s", line)
			continue
		}

		dump := spec.WgDump{
			Interface:     fields[0],
			PublicKey:     fields[1],
			PrivateKey:    fields[2],
			Port:          fields[3],
			Endpoint:      fields[4],
			TransferRx:    fields[5],
			TransferTx:    fields[6],
			LastHandshake: fields[7],
		}

		// Handle optional AllowedIPs and Flags
		if len(fields) > 8 {
			dump.AllowedIPs = fields[8]
		}
		if len(fields) > 9 {
			dump.Flags = fields[9]
		}

		// Append the parsed struct to the list
		result = append(result, dump)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
