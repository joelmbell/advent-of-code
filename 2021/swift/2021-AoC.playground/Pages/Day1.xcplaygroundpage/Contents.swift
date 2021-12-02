//: [Previous](@previous)

import Foundation

let data = try! FileLoader().loadFile(named: "pt1Input")
    .compactMap { Int($0) }

func pt1Solution() {
    func countLarger(input: [Int]) -> Int {
        var previousValue: Int? = nil
        var counter: Int = 0
        for item in input {
            if let prev = previousValue {
                if item > prev {
                    counter += 1    
                }
            }
            previousValue = item
        }
        return counter
    }
    
    print(countLarger(input: data))
}
pt1Solution()


func pt2Solution(input: [Int] = data) -> Int {
    var counter = 0
    var idx = 0
    while idx + 3 < input.count {
        let current = input[idx] + input[idx+1] + input[idx+2]
        let next = input[idx+1] + input[idx+2] + input[idx+3]
        if next > current {
            counter += 1
        }
        idx += 1
    }
    return counter
}

// Testcase
pt2Solution(input: [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]) == 5
pt2Solution()
