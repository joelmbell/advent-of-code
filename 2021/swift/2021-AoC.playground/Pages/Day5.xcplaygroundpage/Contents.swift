//: [Previous](@previous)

import Foundation

let sampleData = try! FileLoader().loadFile(named: "sample")
let data = try! FileLoader().loadFile(named: "input")

struct Point: CustomStringConvertible, Hashable {
    let x: Int
    let y: Int
    
    var description: String {
        return "(\(x), \(y))"
    }
}

struct Line: CustomStringConvertible {
    let start: Point
    let end: Point
    
    init(raw: String) {
        let rawPoints = raw.components(separatedBy: " -> ")
        let points: [Point] = rawPoints.map { rawPoint in
            let pointVals = rawPoint.components(separatedBy: ",")
            return Point(x: Int(pointVals[0])!, y: Int(pointVals[1])!)
        }
        
        self.start = points[0]
        self.end = points[1]
    }
    
    var description: String {
        return "\(start) - \(end)"
    }
    
    var points: [Point] {
        guard start.x == end.x || start.y == end.y else {
            return []
        }
        
        let modifyingX = end.y == start.y
        let changeXBy = end.x - start.x
        let changeYBy = end.y - start.y
        let totalChange = max(abs(changeYBy), abs(changeXBy))
        
        guard totalChange > 0 else {
            return [start]
        }
        
        guard totalChange > 1 else {
            return [start, end]
        }
        
        var points: [Point] = [start]
        for i in 1...totalChange {
            let newPoint: Point
            
            if modifyingX {
                let changeVal = (changeXBy < 0) ? start.x - i : start.x + i
                newPoint = Point(x: changeVal, y: start.y)
            } else {
                let changeVal = (changeYBy < 0) ? start.y - i : start.y + i
                newPoint = Point(x: start.x, y: changeVal)
            }
            
            points.append(newPoint)
        }
        return points
    }
}

func solvePt1(_ data: [String]) -> Int {
    let lines = data
        .filter { !$0.isEmpty }
        .map { Line(raw: $0) }
    
    var map: [Point: Int] = [:]
        
    for line in lines {
        for point in line.points {
            if map[point] == nil {
                map[point] = 1
            } else {
                map[point] = map[point]! + 1
            }
        }
    }
    var count: Int = 0
    for (_, v) in map {
        if v > 1 {
            count += 1
        }
    }
    return count
}

solvePt1(sampleData) == 5
solvePt1(data)

//: [Next](@next)
