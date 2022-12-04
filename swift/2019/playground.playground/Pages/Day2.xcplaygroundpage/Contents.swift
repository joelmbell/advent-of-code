import Foundation

struct FileLoader {
    enum Err: Error {
        case loadError
    }

    func loadFile(named name: String, ext: String = "txt") throws -> [Int] {
        guard
            let filePath = Bundle.main.path(forResource: name, ofType: ext),
            let contentData = FileManager.default.contents(atPath: filePath),
            let contents = String(data: contentData, encoding: .utf8)
            else { throw Err.loadError }

        return contents
            .components(separatedBy: ",")
            .compactMap(Int.init)
    }
}

/**
 Requirements

 opcode 1 -
    adds next two, stores result in 3rd (pointers)
 opcode 2 -
    multiplies next two stores result in 3rd (pointers)
 opcode 99 - end program
 */


struct Interpreter {
    func interprete(_ input: inout [Int]) {

        var i: Int = 0
        while input[i] != 99, i < input.count {
            switch input[i] {
            case 1:
                code1(i, input: &input)
                i += 4
            case 2:
                code2(i, input: &input)
                i += 4
            default: continue
            }
        }
    }

    private func code1(_ i: Int, input: inout [Int]) {
        let result = input[input[i+1]] + input[input[i+2]]
        input[input[i+3]] = result
    }

    private func code2(_ i: Int, input: inout [Int]) {
        let result = input[input[i+1]] * input[input[i+2]]
        input[input[i+3]] = result
    }
}



func solvePart1() -> Int {
    var fileInput = try! FileLoader().loadFile(named: "input")
    fileInput[1] = 12
    fileInput[2] = 2
    Interpreter().interprete(&fileInput)
    return fileInput[0]
 }



func solvePart2() -> Int {
    let interpreter = Interpreter()
    for noun in 0..<100 {
        for verb in 0..<100 {
            var fileInput = try! FileLoader().loadFile(named: "input")
            fileInput[1] = noun
            fileInput[2] = verb
            interpreter.interprete(&fileInput)

            if fileInput[0] == 19690720 {
                return 100 * noun + verb
            }
        }
    }

    return -1
}

print("Part 1 solution: \(solvePart1())")
//print("Part 2 solution: \(solvePart2())")









