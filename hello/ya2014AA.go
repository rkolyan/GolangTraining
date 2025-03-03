package main

type position_t byte;

func make_position_from_input(str_pos string) position_t {
	return position_t(byte(str_pos[0]) - 'a' | (byte(str_pos[1]) - '1') << 4);
}

func min(value1, value2 byte) byte {
	if (value1 < value2) {
		return value1;
	}
	return value2;
}

func max(value1, value2 byte) byte {
	if (value1 > value2) {
		return value1;
	}
	return value2;
}

type distance_t byte;

func get_x_distance(pos1, pos2 position_t) distance_t {
	pos1_x := (pos1 >> 4 & 0b1111);
	pos2_x := (pos2 >> 4 & 0b1111);
	return distance_t(max(pos1_x, pos2_x) - min(pos1_x, pos2_x));
}

func get_y_distance(pos1, pos2 position_t) distance_t {
	pos1_y := (pos1 & 0b1111);
	pos2_y := (pos2 & 0b1111);
	return distance_t(max(pos1_y, pos2_y) - min(pos1_y, pos2_y));
}

func get_distance(pos1, pos2 position_t) distance_t {
	return distance_t(max(distance_t(get_x_distance(pos1, pos2)), distance_t(get_y_distance(pos1, pos2))));
}

func is_edge(pos position_t) bool {
	side_x := (pos >> 4) & 0x0f;
	side_y := pos & 0x0f;
	if (side_x == 0 || side_x == 7) && (side_y == 0 || side_y == 7) {
		return true;
	}
	return false;
}

func found_status(white_king_pos, white_rook_pos, black_king_pos position_t) string {
	if get_distance(white_king_pos, black_king_pos) == 1 {
		return "Strange";
	} else if is_edge(black_king_pos) {
		if get_y_distance(white_rook_pos, black_king_pos) == 0 && get_y_distance(white_king_pos, black_king_pos) == 2 && get_x_distance(white_king_pos, black_king_pos) <= 1 {
			if get_x_distance(white_rook_pos, black_king_pos) > 1 {
				return "Checkmate";
			}
			return "Check";
		} else if get_x_distance(white_rook_pos, black_king_pos) == 0 && get_x_distance(white_king_pos, black_king_pos) == 2 && get_y_distance(white_king_pos, black_king_pos) <= 1 {
			if get_y_distance(white_rook_pos, black_king_pos) > 1 {
				return "Checkmate";
			}
			return "Check";
		} else if (get_x_distance(white_rook_pos, black_king_pos) == 1 && get_y_distance(white_king_pos, black_king_pos) == 2 && get_x_distance(white_king_pos, black_king_pos) <= 1) || (get_y_distance(white_rook_pos, black_king_pos) == 1 && get_x_distance(white_king_pos, black_king_pos) == 2 && get_y_distance(white_king_pos, black_king_pos) <= 1) {
			if (get_distance(white_rook_pos, black_king_pos) == 1) {
				return "Check";
			} else {
				return "Stalemale";
			}
		}
	} else if (get_x_distance(white_rook_pos, black_king_pos) == 0) {
		if (get_x_distance(black_king_pos, white_king_pos) > 0 || get_y_distance(black_king_pos, white_rook_pos) != get_y_distance(black_king_pos, white_king_pos) + get_y_distance(white_king_pos, white_rook_pos)) {
			return "Check";
		}
	} else if (get_y_distance(white_rook_pos, black_king_pos) == 0) {
		if (get_y_distance(black_king_pos, white_king_pos) > 0 || get_x_distance(black_king_pos, white_rook_pos) != get_x_distance(black_king_pos, white_king_pos) + get_x_distance(white_king_pos, white_rook_pos)) {
			return "Check";
		}
	}
	return "Normal";
}

/*

func main() {
	white_king := "b2";
	white_rook := "a3";
	black_king := "c4";
	white_king_pos := make_position_from_input(white_king);
	white_rook_pos := make_position_from_input(white_rook);
	black_king_pos := make_position_from_input(black_king);
	fmt.Println("Status = ", found_status(white_king_pos, white_rook_pos, black_king_pos));
}

*/
