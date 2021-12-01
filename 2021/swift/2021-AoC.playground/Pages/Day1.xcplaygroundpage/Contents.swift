//: [Previous](@previous)

import Foundation


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

func pt1Solution() {
    let data = try! FileLoader().loadFile(named: "pt1Input")
        .compactMap { Int($0) }
    print(countLarger(input: data))
}
pt1Solution()

//: [Next](@next)
