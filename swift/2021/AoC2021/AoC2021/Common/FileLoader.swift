import Foundation

struct FileLoader {
    enum Err: Error {
        case loadError
    }
    
    func loadFile(at path: String) throws -> [String] {
        let fullPath = "/Users/i508274/dev/advent-of-code/2021/swift/AoC2021/AoC2021/" + path
        guard
            let contentData = FileManager.default.contents(atPath: fullPath),
            let contents = String(data: contentData, encoding: .utf8)
            else { throw Err.loadError }

        return contents.components(separatedBy: "\n")
    }
}
