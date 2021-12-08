//
//  Day7.swift
//  AoC2021
//
//  Created by Bell, Joel on 12/7/21.
//

import Foundation

private func solvePt1(_ data: [String]) -> Int {
    let input = data
        .filter { !$0.isEmpty }
        .first!
        .components(separatedBy: ",")
        .map { Int($0)! }
    
    var cheapest: (hPos: Int, fuelCost: Int)? = nil
    for i in 0..<input.count {
        var cost: Int = 0
        for k in 0..<input.count {
            let diff = abs(input[i] - input[k])
            cost += diff
            if let cheapest = cheapest, cost > cheapest.fuelCost {
                break
            }
        }
        
        if cheapest == nil {
            cheapest = (input[i], cost)
        } else {
            if cost < cheapest!.fuelCost {
                cheapest = (input[i], cost)
            }
        }
    }
    
    return cheapest!.fuelCost
}

private func solvePt2(_ data: [String]) -> Int {
    let input = data
        .filter { !$0.isEmpty }
        .first!
        .components(separatedBy: ",")
        .map { Int($0)! }
        .sorted(by: >)
        
    var cheapest: (hPos: Int, fuelCost: Int)? = nil
    for i in 0..<input[0] {
        var cost: Int = 0
        for k in 0..<input.count {
            let diff = abs(i - input[k])
            var fuelCost = 0
            for f in 0..<diff {
                fuelCost += 1+f
            }
            
            cost += fuelCost
            if let cheapest = cheapest, cost > cheapest.fuelCost {
                break
            }
        }

        if cheapest == nil {
            cheapest = (i, cost)
        } else {
            if cost < cheapest!.fuelCost {
                cheapest = (i, cost)
            }
        }
    }
    
    return cheapest!.fuelCost
}

func runDay7() {
    let sampleData = try! FileLoader().loadFile(at: "Day7/sample.txt")
    let data = try! FileLoader().loadFile(at: "Day7/input.txt")
    
    print("pt 1 sample: \(solvePt1(sampleData) == 37)")
    print("pt 1 data: \(solvePt1(data) == 336120)")

    print("pt 2 sample: \(solvePt2(sampleData) == 168)")
    print("pt 2 data: \(solvePt2(data) == 96864235)")
}

