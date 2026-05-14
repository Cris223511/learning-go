// Acá veremos pgx directamente: más eficiente que database/sql para PostgreSQL.
// pgxpool maneja el pool de conexiones y pgx.CollectRows simplifica el escaneo.
//
// Requisito: PostgreSQL corriendo (ver instrucciones en 01-database-sql).
// Correr: go run ./11-postgresql/02-pgx

package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	pool, err := NuevoPool(ctx)
	if err != nil {
		log.Fatal("no se pudo conectar:", err)
	}
	defer pool.Close()
	fmt.Println("pool de conexiones listo")

	repo := NuevoPolizaRepo(pool)

	if err := repo.CrearTabla(ctx); err != nil {
		log.Fatal("error creando tabla:", err)
	}

	fmt.Println("\n== Insertar pólizas ==")
	ids := []int{}
	polizas := []Poliza{
		{ClienteID: 1, Tipo: "SOAT", Prima: 120.00, Activa: true},
		{ClienteID: 1, Tipo: "Vida", Prima: 1500.00, Activa: true},
		{ClienteID: 2, Tipo: "Vehicular", Prima: 890.75, Activa: false},
	}
	for _, p := range polizas {
		id, err := repo.Insertar(ctx, p)
		if err != nil {
			log.Println("error insertando:", err)
			continue
		}
		ids = append(ids, id)
		fmt.Printf("  insertada póliza ID=%d | %s | S/%.2f\n", id, p.Tipo, p.Prima)
	}

	fmt.Println("\n== Listar todas ==")
	todas, err := repo.Listar(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range todas {
		estado := "activa"
		if !p.Activa {
			estado = "inactiva"
		}
		fmt.Printf("  [%d] %s | S/%.2f | %s\n", p.ID, p.Tipo, p.Prima, estado)
	}

	if len(ids) > 0 {
		fmt.Println("\n== Buscar por ID ==")
		p, err := repo.BuscarPorID(ctx, ids[0])
		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Printf("  encontrada: %+v\n", p)
		}

		fmt.Println("\n== Actualizar ==")
		if err := repo.Actualizar(ctx, ids[0], false); err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Printf("  póliza %d desactivada\n", ids[0])
		}
	}
}
