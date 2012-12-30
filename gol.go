package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type gol struct {
	world     []byte
	new_world []byte
	sides     [][8]int
	born      []int
	alive     []int
}

func get_rules(rules string) ([]int, []int) {
	r := strings.Split(rules, "/")
	born := make([]int, len(r[0][1:]))
	alive := make([]int, len(r[1][1:]))
	for i, e := range r[0][1:] {
		born[i] = int(e - 48)
	}
	for i, e := range r[1][1:] {
		alive[i] = int(e - 48)
	}
	return born, alive
}

func side_cells(C, L int) [][8]int {
	sides := make([][8]int, C*L)
	for i := 0; i < L; i++ {
		for j := 0; j < C; j++ {
			sides[i*C+j][0] = ((((i-2)%L + 1 + L) * C) + (((j-2+C)%C + 1) % C)) % (C * L)
			sides[i*C+j][1] = (((i-2)%L+1+L)*C + j) % (C * L)
			sides[i*C+j][2] = ((((i-2)%L + 1 + L) * C) + (((j+C)%C + 1) % C)) % (C * L)
			sides[i*C+j][3] = (i*C + ((j-2+C)%C+1)%C) % (C * L)
			sides[i*C+j][4] = (i*C + ((j+C)%C+1)%C) % (C * L)
			sides[i*C+j][5] = (((i%L + 1 + L) * C) + (((j-2+C)%C + 1) % C)) % (C * L)
			sides[i*C+j][6] = ((i%L+1+L)*C + j) % (C * L)
			sides[i*C+j][7] = (((i%L + 1 + L) * C) + (((j+C)%C + 1) % C)) % (C * L)
		}
	}
	return sides
}

func count_alive(world []byte, sides [8]int) int {
	alive := 0
	for i := 0; i < 8; i++ {
		if world[sides[i]] == 1 {
			alive++
		}
	}
	return alive
}

func is_in(value int, data []int) bool {
	in := false
	for i := range data {
		if data[i] == value {
			in = true
			break
		}
	}
	return in
}

func cycle(data gol) []byte {
	for pos := range data.world {
		alive := count_alive(data.world, data.sides[pos])
		if data.world[pos] == 1 {
			if is_in(alive, data.alive) {
				data.new_world[pos] = 1
			} else {
				data.new_world[pos] = 0
			}
		} else {
			if is_in(alive, data.born) {
				data.new_world[pos] = 1
			} else {
				data.new_world[pos] = 0
			}
		}
	}
	return data.new_world
}

func dump(world []byte, C, L int) {
	fmt.Println("--------------------------------------------------------------------------------")
	for i := 0; i < (C * L); i += C {
		fmt.Println(world[i : i+C])
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func import_map(m string, c, l int, world []byte) []byte {
	if len(m) != 0 {
		for i := range world {
			world[i] = m[i] - 48
		}
	}
	return world
}

func import_file(f string, world []byte) (error, []byte) {
	if len(f) != 0 {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return err, world
		}
		pos := 0
		for _, c := range string(b) {
			if c == '1' || c == '0' {
				world[pos] = byte(c) - 48
				pos++
			}
		}
	}
	return nil, world
}

func Init(rules, m, f string, c, l, cycle int) (error, gol) {
	born, alive := get_rules(rules)
	world := make([]byte, c*l)
	new_world := make([]byte, c*l)
	sides := side_cells(c, l)
	world = import_map(m, c, l, world)
	err, world := import_file(f, world)
	game := gol{world, new_world, sides, born, alive}
	return err, game
}

func Launch(columns, lines, cycles int, data gol) {
	dump(data.world, columns, lines)
	for i := 0; i < cycles; i++ {
		data.new_world = cycle(data)
		copy(data.world, data.new_world)
		dump(data.world, columns, lines)
	}
}
