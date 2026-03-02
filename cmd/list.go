package cmd

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"pdd/store"
)

func List() {
	puzzles, _ := store.Load()

	// Сортировка: сначала OPEN, внутри — по времени (новые внизу)
	sort.Slice(puzzles.Items, func(i, j int) bool {
		if puzzles.Items[i].Status == puzzles.Items[j].Status {
			return puzzles.Items[i].Timestamp < puzzles.Items[j].Timestamp
		}
		return puzzles.Items[i].Status == "OPEN"
	})

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tSTATUS\tTYPE\tFILE\tDESCRIPTION")
	fmt.Fprintln(w, "--\t------\t----\t----\t-----------")

	for _, p := range puzzles.Items {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
			p.Timestamp,
			p.Status,
			p.Type,
			p.File,
			p.Description,
		)
	}
	w.Flush()
}
