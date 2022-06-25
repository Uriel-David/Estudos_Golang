package main

import "fmt"

func main() {
	x := `
		<script type="text/javascript">
			alert("Será que funfa?");
		</script>
	`

	fmt.Println(x)
}
