//
//  Day5.swift
//  AoC2021
//
//  Created by Bell, Joel on 12/7/21.
//

import Foundation

private struct Point: CustomStringConvertible, Hashable {
    let x: Int
    let y: Int
    
    var description: String {
        return "(\(x), \(y))"
    }
}

private struct Line: CustomStringConvertible {
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
    
    var isHorizontal: Bool {
        return (end.x - start.x) == 0
    }
    
    var isVertical: Bool {
        return (end.y - start.y) == 0
    }
    
    var description: String {
        return "\(start) - \(end)"
    }
            
    func drawLine(into map: inout [Point: Int]) {
        let changeXBy = end.x - start.x
        let changeYBy = end.y - start.y
        let totalChange = max(abs(changeYBy), abs(changeXBy))
        
        insert(point: start, into: &map)
        insert(point: end, into: &map)
        
        for i in 1..<totalChange {
            var newX: Int
            var newY: Int
            
            if changeXBy == 0 {
                newX = start.x
            } else {
                newX = changeXBy < 0 ? start.x - i : start.x + i
            }
            
            if changeYBy == 0 {
                newY = start.y
            } else {
                newY = changeYBy < 0 ? start.y - i : start.y + i
            }
            
            let newPoint = Point(x: newX, y: newY)
            insert(point: newPoint, into: &map)
        }
    }
    
    private func insert(point: Point, into map: inout [Point: Int]) {
        let current = map[point] ?? 0
        map[point] = current + 1
    }
}

private func solveDay5Pt1(_ data: [String]) -> Int {
    let lines = data
        .filter { !$0.isEmpty }
        .map { Line(raw: $0) }
    
    return lines
        .reduce(into: [Point: Int]()) { partialResult, line in
            guard line.isHorizontal || line.isVertical else { return }
            line.drawLine(into: &partialResult)
        }
        .reduce(into: 0) { partialResult, item in
            if item.value > 1 {
                partialResult += 1
            }
        }
}

private func solveDay5Pt2(_ data: [String]) -> Int {
    let lines = data
        .filter { !$0.isEmpty }
        .map { Line(raw: $0) }
    
    return lines
        .reduce(into: [Point: Int]()) { partialResult, line in
            line.drawLine(into: &partialResult)
        }
        .reduce(into: 0) { partialResult, item in
            if item.value > 1 {
                partialResult += 1
            }
        }
}

func runDay5() {
    let sampleData = try! FileLoader().loadFile(at: "Day5/sample.txt")
    let data = try! FileLoader().loadFile(at: "Day5/input.txt")
    
    print("pt 1 sample data: \(solveDay5Pt1(sampleData) == 5)")
    print("pt 1 data: \(solveDay5Pt1(data) == 5124)")
    
    print("pt 2 sample data: \(solveDay5Pt2(sampleData) == 12)")
    print("pt 2 data: \(solveDay5Pt2(data) == 19771)")
}

