//: [Previous](@previous)

import Foundation

let sampleData = try! FileLoader().loadFile(named: "sample")
let data = try! FileLoader().loadFile(named: "input")

struct Board: CustomStringConvertible {
    struct Point {
        let number: Int
        var isMarked: Bool
        
        init(number: Int, isMarked: Bool = false) {
            self.number = number
            self.isMarked = isMarked
        }
    }
    
    private(set) var data: [[Point]] = []
    
    init(rawData: [String]) {
        data = rawData.map { 
            $0.split(separator: " ").map { Point(number: Int($0)!) }
        }
    }
    
    
    var isWinner: Bool {
        
        func isWinner(_ points: [Point]) -> Bool {
            return points.filter { $0.isMarked }.count == points.count
        }
        
        var rotatedData: [[Point]] = Array(repeating: [Point](), count: data[0].count)
        for i in 0..<data.count {
            for k in 0..<data[i].count {
                rotatedData[i].append(data[k][i])
            }
        }
        
        // Check Rows & Columns
        for i in 0..<data.count {
            if isWinner(data[i]) || isWinner(rotatedData[i]) {
                return true
            }
        }
               
        return false
    }
    
    mutating func play(num: Int) {
        for i in 0..<data.count {
            for k in 0..<data[i].count {
                if data[i][k].number == num {
                    data[i][k].isMarked = true
                }
            }
        }
    }
    
    var description: String {
        var output: String = "\n"
        for i in 0..<data.count {
            output += "\n"
            for k in 0..<data[i].count {
                let point = data[i][k]
                let symbol = point.isMarked ? "+" : "-"
                output += "\(point.number)\(symbol) "
            }
            output += "\n"
        }
        output += "\n"
        
        return output
    }
}

extension Board {
    var sumUnmarkedNumbers: Int {
        return self.data.reduce(into: 0) { partialResult, points in
            let rowSum = points.reduce(into: 0) { pointsResult, point in
                if !point.isMarked {
                    pointsResult += point.number
                }
            }
            partialResult += rowSum
        }
    }
}


func solveDay4Part1(data: [String]) -> Int {
    
    let plays = data.first!.split(separator: ",").map { Int($0)! }
    
    var boards: [Board] = []
    var currentBoardData: [String] = []
    for i in 2..<data.count {
        let line = data[i]
        if line.isEmpty, !currentBoardData.isEmpty {
            boards.append(Board(rawData: currentBoardData))
            currentBoardData = []
        } else {
            currentBoardData.append(line)    
        }
    }
    
    for play in plays {
        for i in 0..<boards.count {
            boards[i].play(num: play)
            if boards[i].isWinner {
                return boards[i].sumUnmarkedNumbers * play
            }
        }
    }
        
    fatalError("No winner found")
}

Assert(solveDay4Part1(data: sampleData), 4512)
Assert(solveDay4Part1(data: data), 89001)

//: [Next](@next)
