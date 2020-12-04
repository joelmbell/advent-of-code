import Foundation

let data = try! FileLoader()
    .loadFile(named: "input")
    .filter { !$0.isEmpty }

typealias Slope = (x: Int, y: Int)
func solve(slope: Slope) -> Int {
    var xPos = 0
    var treesEncountered = 0
    for i in stride(from: 0, to: data.count - slope.y, by: slope.y) {
        let currentRow = data[i] as NSString
        let nextRow = data[i+slope.y] as NSString
                
        if currentRow.length <= xPos + slope.x {
            let diff = xPos + slope.x - currentRow.length
            xPos = diff
        } else {
            xPos += slope.x
        }
        
        let char = Character(nextRow.character(at: xPos))
        if char == Character("#") {
            treesEncountered += 1
        }
    }

    return treesEncountered
}

func solvePt2() -> Int {
    [(1, 1),
     (3, 1),
     (5, 1),
     (7, 1),
     (1, 2)]
        .map { solve(slope: $0) }
        .reduce(1, *)
}

func solvePt1() -> Int {
    return solve(slope: (3, 1))
}



print(solvePt1())
//print(solvePt2())




