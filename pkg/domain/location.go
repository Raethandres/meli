package domain

import . "meli/cmd/data"

func GetLocation(distances ...float32) (x, y float32) {
	if len(distances) != 3 {
		return 0, 0
	}

	radiusKenobi := distances[0]
	radiusSkywalker := distances[1]
	rediusSato := distances[2]

	S := (square(SATO_X) - square(SKYWALKER_X) + square(SATO_Y) - square(SKYWALKER_Y) + square(radiusSkywalker) - square(rediusSato)) / 2.0
	T := (square(KENOBI_X) - square(SKYWALKER_X) + square(KENOBI_Y) - square(SKYWALKER_Y) + square(radiusSkywalker) - square(radiusKenobi)) / 2.0

	y = ((T * (SKYWALKER_X - SATO_X)) - (S * (SKYWALKER_X - KENOBI_X))) / (((KENOBI_Y - SKYWALKER_Y) * (SKYWALKER_X - SATO_X)) - ((SATO_Y - SKYWALKER_Y) * (SKYWALKER_X - KENOBI_X)))
	x = ((y * (KENOBI_Y - SKYWALKER_Y)) - T) / (SKYWALKER_X - KENOBI_X)

	return x, y
}

func square(n float32) float32 {
	return n * n
}
