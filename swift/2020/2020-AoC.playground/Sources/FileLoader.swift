import Foundation

public struct FileLoader {
    public enum Err: Error {
        case loadError
    }
    
    public init() {}

    public func loadFile(named name: String, ext: String = "txt") throws -> [String] {
        guard
            let filePath = Bundle.main.path(forResource: name, ofType: ext),
            let contentData = FileManager.default.contents(atPath: filePath),
            let contents = String(data: contentData, encoding: .utf8)
            else { throw Err.loadError }

        return contents.components(separatedBy: "\n")
    }
}

public extension Character {
    public init?(_ uint16: UInt16) {
        guard let scalar = Unicode.Scalar(uint16) else {
            return nil
        }
        self.init(scalar)
    }
}
