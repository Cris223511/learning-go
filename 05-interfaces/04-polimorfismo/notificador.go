// Tipos que implementan la interface Notificador. Cada uno tiene su propia
// lógica de envío pero desde afuera se usan de forma idéntica.

package main

import "fmt"

type Notificador interface {
	Enviar(destinatario, mensaje string) error
	Nombre() string
}

type NotificadorEmail struct {
	Servidor string
	Puerto   int
}

func (n NotificadorEmail) Nombre() string { return "Email" }

func (n NotificadorEmail) Enviar(destinatario, mensaje string) error {
	fmt.Printf("  [Email via %s:%d] para %s: %q\n", n.Servidor, n.Puerto, destinatario, mensaje)
	return nil
}

type NotificadorSMS struct {
	Proveedor string
}

func (n NotificadorSMS) Nombre() string { return "SMS" }

func (n NotificadorSMS) Enviar(destinatario, mensaje string) error {
	if len(mensaje) > 160 {
		return fmt.Errorf("SMS supera los 160 caracteres")
	}
	fmt.Printf("  [SMS via %s] para %s: %q\n", n.Proveedor, destinatario, mensaje)
	return nil
}

type NotificadorPush struct {
	AppID string
}

func (n NotificadorPush) Nombre() string { return "Push" }

func (n NotificadorPush) Enviar(destinatario, mensaje string) error {
	fmt.Printf("  [Push app:%s] para %s: %q\n", n.AppID, destinatario, mensaje)
	return nil
}
