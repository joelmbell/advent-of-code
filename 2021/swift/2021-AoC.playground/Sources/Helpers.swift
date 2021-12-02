import Foundation

public func Assert<T: Equatable>(_ value1: T, _ value2: T) -> Bool {
    guard value1 == value2 else {
        fatalError("\(value1) != \(value2)")
    }
    return true
}

public extension Character {
    init?(_ uint16: UInt16) {
        guard let scalar = Unicode.Scalar(uint16) else {
            return nil
        }
        self.init(scalar)
    }
}
