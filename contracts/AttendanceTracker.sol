// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract AttendanceTracker {
    struct AttendanceRecord {
        uint employeeId;
        uint256 checkInTime;
        string details;
    }

    mapping(uint => AttendanceRecord) public attendanceRecords;
    mapping(uint => uint[]) public employeeAttendanceLog;

    address public owner;

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Unauthorized");
        _;
    }

    event Debug(uint indexed employeeId, uint256 checkInTime, string details);

    function recordAttendance(
        uint employeeId,
        uint256 checkInTime,
        string memory details
    ) public onlyOwner {
        emit Debug(employeeId, checkInTime, details);
        AttendanceRecord memory newRecord = AttendanceRecord(
            employeeId,
            checkInTime,
            details
        );
        attendanceRecords[employeeId] = newRecord;
        employeeAttendanceLog[employeeId].push(checkInTime);
    }

    function getAttendanceByEmployee(
        uint employeeId
    ) public view returns (AttendanceRecord memory) {
        return attendanceRecords[employeeId];
    }

    function updateAttendance(
        uint employeeId,
        uint256 checkInTime,
        string memory details
    ) public onlyOwner {
        require(
            attendanceRecords[employeeId].checkInTime != 0,
            "Record does not exist."
        );
        emit Debug(employeeId, checkInTime, details);
        attendanceRecords[employeeId] = AttendanceRecord(
            employeeId,
            checkInTime,
            details
        );
    }

    function getAttendanceByDateRange(
        uint employeeId,
        uint256 startTime,
        uint256 endTime
    ) public view returns (AttendanceRecord[] memory) {
        uint[] memory log = employeeAttendanceLog[employeeId];
        AttendanceRecord[] memory recordsInRange = new AttendanceRecord[](
            log.length
        );

        uint count = 0;
        for (uint i = 0; i < log.length; i++) {
            AttendanceRecord memory record = attendanceRecords[log[i]];
            if (
                record.checkInTime >= startTime && record.checkInTime <= endTime
            ) {
                recordsInRange[count] = record;
                count++;
            }
        }

        // Cắt mảng theo số lượng bản ghi thực tế được tìm thấy
        AttendanceRecord[] memory trimmedRecords = new AttendanceRecord[](
            count
        );
        for (uint j = 0; j < count; j++) {
            trimmedRecords[j] = recordsInRange[j];
        }

        return trimmedRecords;
    }
}
