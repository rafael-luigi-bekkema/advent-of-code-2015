const print = @import("std").debug.print;
const std = @import("std");
const assert = @import("std").debug.assert;
const expect = @import("std").testing.expect;

test "day1" {
    try expect(day1("(())") == 0);
    try expect(day1("()()") == 0);
    try expect(day1("(()(()(") == 3);
    try expect(day1(")())())") == -3);

    try expect(day1(try day1file()) == 232);
}

test "day1b" {
    try expect(day1b(")") == 1);
    try expect(day1b("()())") == 5);

    try expect(day1b(try day1file()) == 1783);
}

fn day1file() ![]const u8 {
    const file = try std.fs.cwd().openFile(
        "input/day01.txt",
        .{
            .read = true,
            .write = false,
        },
    );
    defer file.close();

    var buf: [1024*7]u8 = undefined;
    const l = try file.reader().readUntilDelimiterOrEof(&buf, '\n');
    if (l) |line| {
        return line;
    }
    return "";
}

fn day1b(input: []const u8) i64 {
    var n: i64 = 0;
    var pos: i64 = 0;
    for (input) |ch| {
        pos+=1;
        if (ch == '(') {
            n+=1;
        } else {
            n-=1;
        }
        if (n < 0) {
            return pos;
        }
    }
    return -1;
}

fn day1(input: []const u8) i64 {
    var n: i64 = 0;
    for (input) |ch| {
        if (ch == '(') {
            n+=1;
        } else {
            n-=1;
        }
    }
    return n;
}

pub fn main() void {
    print("hello guy", .{});
}

