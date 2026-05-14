// Acá veremos polimorfismo al estilo Go: una función que trabaja con una interface
// y se comporta diferente según el tipo concreto que reciba, sin saber cuál es.

package main

import "fmt"

// Esta función no sabe si está enviando un email, SMS o push.
// Solo sabe que lo que recibe cumple Notificador. Eso es polimorfismo.
func notificarPoliza(n Notificador, cliente, polizaID string) {
	msg := fmt.Sprintf("Tu póliza %s ha sido activada.", polizaID)
	if err := n.Enviar(cliente, msg); err != nil {
		fmt.Printf("  [%s] error: %v\n", n.Nombre(), err)
	}
}

// Puedes combinar varias implementaciones y usarlas todas con el mismo código.
func notificarTodos(notificadores []Notificador, cliente, mensaje string) {
	fmt.Printf("notificando a %s por %d canal(es):\n", cliente, len(notificadores))
	for _, n := range notificadores {
		if err := n.Enviar(cliente, mensaje); err != nil {
			fmt.Printf("  [%s] falló: %v\n", n.Nombre(), err)
		}
	}
}

// fmt.Stringer: si tu tipo implementa String() string, fmt.Println lo usa automáticamente.
// Es la interface más usada de la stdlib junto con error.
type Producto struct {
	Codigo string
	Nombre string
	Prima  float64
}

func (p Producto) String() string {
	return fmt.Sprintf("Producto{%s | %s | S/%.2f}", p.Codigo, p.Nombre, p.Prima)
}

func main() {
	email := NotificadorEmail{Servidor: "smtp.correo.com", Puerto: 587}
	sms := NotificadorSMS{Proveedor: "Twilio"}
	push := NotificadorPush{AppID: "app.miapi.com"}

	fmt.Println("== Polimorfismo: misma función, distintos comportamientos ==")
	notificarPoliza(email, "christopher@mail.com", "POL-001")
	notificarPoliza(sms, "+51999888777", "POL-002")
	notificarPoliza(push, "device-token-abc123", "POL-003")

	fmt.Println("\n== Slice de interfaces: múltiples canales ==")
	canales := []Notificador{email, sms, push}
	notificarTodos(canales, "cliente@mail.com", "Tu póliza vence en 7 días.")

	fmt.Println("\n== Error controlado (SMS demasiado largo) ==")
	largo := "Este es un mensaje extremadamente largo que supera el límite de ciento sesenta caracteres permitidos por el protocolo SMS estándar internacional vigente."
	notificarPoliza(sms, "+51999888777", largo)

	fmt.Println("\n== fmt.Stringer ==")
	p := Producto{Codigo: "SEG-SOAT", Nombre: "SOAT Nacional", Prima: 120.00}
	fmt.Println(p)        // Go llama a String() automáticamente
	fmt.Printf("%v\n", p) // %v también usa String()
}
