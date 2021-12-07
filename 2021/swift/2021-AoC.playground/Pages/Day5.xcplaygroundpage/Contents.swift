//: [Previous](@previous)

import Foundation

let sampleData = try! FileLoader().loadFile(named: "sample")
let data = try! FileLoader().loadFile(named: "input")

struct Point: CustomStringConvertible {
    let x: Int
    let y: Int
    
    var description: String {
        return "(\(x), \(y))"
    }
    
    var uniqueId: String {
        return "\(x),\(y)"
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
        insertTime = 0
    }
    
    var description: String {
        return "\(start) - \(end)"
    }
    
    var insertTime: CFAbsoluteTime
    
    mutating func drawLine(into map: inout [String: Int]) {
        func insert(point: Point, into map: inout [String: Int]) {
            var timer = ParkBenchTimer()
            map[point.uniqueId] = (map[point.uniqueId] == nil) ? 1 : map[point.uniqueId]! + 1
            insertTime += timer.stop()
        }
        
        let modifyingX = end.y == start.y
        let changeXBy = end.x - start.x
        let changeYBy = end.y - start.y
        let totalChange = max(abs(changeYBy), abs(changeXBy))
        
        guard totalChange > 0 else {
            insert(point: start, into: &map)
            return
        }
        
        guard totalChange > 1 else {
            insert(point: start, into: &map)
            insert(point: end, into: &map)
            return
        }
        
        for i in 0...totalChange {
            let newPoint: Point
            
            if modifyingX {
                let changeVal = (changeXBy < 0) ? start.x - i : start.x + i
                newPoint = Point(x: changeVal, y: start.y)
            } else {
                let changeVal = (changeYBy < 0) ? start.y - i : start.y + i
                newPoint = Point(x: start.x, y: changeVal)
            }
            insert(point: newPoint, into: &map)
        }
    }
}

func solvePt1(_ data: [String]) -> Int {
    var timer = ParkBenchTimer()
    let lines = data
        .filter { !$0.isEmpty }
        .map { Line(raw: $0) }
    
    var map: [String: Int] = [:]
        
    for line in lines {
        var line = line
        guard line.start.x == line.end.x || line.start.y == line.end.y else {
            continue
        }
        line.drawLine(into: &map)
    }
    var count: Int = 0
    for (_, v) in map {
        if v > 1 {
            count += 1
        }
    }
    print("insert duration: \(timer.stop())")
    print("duration: \(timer.stop())")
    return count
}

func solvePt2(_ data: [String]) -> Int {
    var timer = ParkBenchTimer()
    let lines = data
        .filter { !$0.isEmpty }
        .map { Line(raw: $0) }
    
    var map: [String: Int] = [:]
        
    for line in lines {
        var line = line
        line.drawLine(into: &map)
    }
    var count: Int = 0
    for (_, v) in map {
        if v > 1 {
            count += 1
        }
    }
    print("duration: \(timer.stop())")
    return count
}

solvePt1(sampleData) == 5
//solvePt1(data) == 5124
//solvePt2(sampleData) == 12

//: [Next](@next)
