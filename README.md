# Lem-in

Welcome to the **lem-in** project! This Go-based program simulates an ant colony, where ants must traverse from a starting room to an exit in the quickest and most efficient way possible, avoiding traffic jams and obeying constraints. The project is designed to improve algorithmic thinking and problem-solving skills.

## Overview

The objective of the program is to:
1. Parse a colony description from a file.
2. Determine the quickest path for ants to travel from the start to the end.
3. Display the movement of ants through the rooms in a formatted output.

## Features

- Create and manage ant farms with rooms and tunnels.
- Find optimal paths for ants to traverse.
- Detect and handle invalid input formats gracefully.
- Simulate ant movements and display them step-by-step.

## Input Format

The program accepts a file describing:
1. **Number of ants**.
2. **Rooms**: Each room is defined as `name coord_x coord_y`. Special commands:
   - `##start` specifies the starting room.
   - `##end` specifies the ending room.
3. **Links**: Defined as `room1-room2`.

### Example Input

```txt
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1
```

## Output Format

- The program outputs the input followed by the movement of ants. Each movement is in the format `Lx-y`, where:
  - `x` is the ant number.
  - `y` is the room name.

### Example Output

```txt
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3
L3-1
```

## Constraints

- A room name must not start with `L` or `#` and must contain no spaces.
- Each tunnel connects exactly two rooms.
- Ants cannot occupy the same room (except `##start` and `##end`).
- Tunnels can only be used once per turn.
- All coordinates must be integers.

## Error Handling

The program gracefully handles invalid inputs with messages such as:
- `ERROR: invalid data format`
- `ERROR: no start room found`
- `ERROR: invalid number of ants`

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Vincent-Omondi/lem-in.git
   cd lem-in
   ```
2. Ensure you have Go installed.
3. Build or run the project:
   ```bash
   go run . <input_file>
   ```

## Usage

### Example Usage

1. Save your colony description to a file (e.g., `test0.txt`).
2. Run the program:
   ```bash
   go run . test0.txt
   ```

### Sample Colony and Output

**Input File:**

```txt
3
##start
0 1 0
##end
1 5 0
2 9 0
0-1
1-2
```

**Output:**

```txt
3
##start
0 1 0
##end
1 5 0
2 9 0
0-1
1-2

L1-1
L1-2 L2-1
L2-2
```

## Development

This project is implemented in Go and adheres to good coding practices. Contributions are welcome!


### Testing

Unit tests are provided for robust validation of functionality. Run tests using:
```bash
go test ./tests
```

## Future Enhancements

- Add a graphical representation of the ant movements.
- Extend error handling to provide more specific feedback.
- Optimize the pathfinding algorithm for larger colonies.

---

Contribute to the project or report issues on GitHub. Let’s build the most efficient ant colony together! 

[Benard](https://github.com/bernotieno) | [Vincent](https://github.com/Vincent-Omondi/lem-in) | [Stella](https://github.com/Stella-Achar-Oiro)