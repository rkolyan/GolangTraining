package main

import "fmt"

func main() {

	var white_king, white_rook, black_king string;
	fmt.Scanf("%s %s %s", &white_king, &white_rook, &black_king);

	white_king_pos := make_position_from_input(white_king);
	white_rook_pos := make_position_from_input(white_rook);
	black_king_pos := make_position_from_input(black_king);
	fmt.Println("Status = ", found_status(white_king_pos, white_rook_pos, black_king_pos));
}
