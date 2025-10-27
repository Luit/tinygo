//go:build tiny2040

package machine

// GPIO pins
const (
	GP0 Pin = GPIO0
	GP1 Pin = GPIO1
	GP2 Pin = GPIO2
	GP3 Pin = GPIO3
	GP4 Pin = GPIO4
	GP5 Pin = GPIO5
	GP6 Pin = GPIO6
	GP7 Pin = GPIO7

	A0 Pin = GPIO26
	A1 Pin = GPIO27
	A2 Pin = GPIO28
	A3 Pin = GPIO29

	// Onboard LED
	LED_R Pin = GPIO18
	LED_G Pin = GPIO19
	LED_B Pin = GPIO20

	BUTTON Pin = GPIO23

	// Onboard crystal oscillator frequency, in MHz.
	xoscFreq = 12 // MHz
)

// I2C Default pins on Raspberry Pico.
const (
	I2C0_SDA_PIN = GP4
	I2C0_SCL_PIN = GP5

	I2C1_SDA_PIN = GP2
	I2C1_SCL_PIN = GP3
)

// SPI default pins
const (
	// Default Serial Clock Bus 0 for SPI communications
	SPI0_SCK_PIN = GP2
	// Default Serial Out Bus 0 for SPI communications
	SPI0_SDO_PIN = GP3 // Tx
	// Default Serial In Bus 0 for SPI communications
	SPI0_SDI_PIN = GP0 // Rx

	// SPI1 not exposed to external pins, but configured like a Pico to please
	// machine_rp2_spi.go

	// Default Serial Clock Bus 1 for SPI communications
	SPI1_SCK_PIN = GPIO10
	// Default Serial Out Bus 1 for SPI communications
	SPI1_SDO_PIN = GPIO11 // Tx
	// Default Serial In Bus 1 for SPI communications
	SPI1_SDI_PIN = GPIO12 // Rx

)

// UART pins
const (
	UART0_TX_PIN = GP0
	UART0_RX_PIN = GP1
	UART1_TX_PIN = GP4
	UART1_RX_PIN = GP5
	UART_TX_PIN  = UART0_TX_PIN
	UART_RX_PIN  = UART0_RX_PIN
)

var DefaultUART = UART0

// USB identifiers
const (
	usb_STRING_PRODUCT      = "Pico"
	usb_STRING_MANUFACTURER = "Raspberry Pi"
)

var (
	usb_VID uint16 = 0x2E8A
	usb_PID uint16 = 0x000A
)
