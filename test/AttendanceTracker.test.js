const AttendanceTracker = artifacts.require("AttendanceTracker");

contract("AttendanceTracker", accounts => {
    let tracker;

    before(async () => {
        tracker = await AttendanceTracker.deployed();
    });

    it("should set the owner correctly", async () => {
        const owner = await tracker.owner();
        assert.equal(owner, accounts[0], "Owner is not set to the first account");
    });

    it("should record attendance correctly", async () => {
        const employeeId = 1;
        const checkInTime = Math.floor(Date.now() / 1000);
        const details = "Employee checked in";

        await tracker.recordAttendance(employeeId, checkInTime, details, { from: accounts[0] });

        const record = await tracker.getAttendanceByEmployee(employeeId);
        assert.equal(record.employeeId.toNumber(), employeeId, "Employee ID does not match");
        assert.equal(record.checkInTime.toNumber(), checkInTime, "Check-in time does not match");
        assert.equal(record.details, details, "Details do not match");
    });

    it("should only allow the owner to record attendance", async () => {
        try {
            await tracker.recordAttendance(2, Math.floor(Date.now() / 1000), "Unauthorized check-in", { from: accounts[1] });
            assert.fail("The transaction should have thrown an error");
        } catch (err) {
            assert.include(err.message, "revert", "The error message should contain 'revert'");
        }
    });

    it("should update attendance correctly", async () => {
        const employeeId = 1;
        const initialCheckInTime = Math.floor(Date.now() / 1000);
        const initialDetails = "Initial check-in";
        await tracker.recordAttendance(employeeId, initialCheckInTime, initialDetails, { from: accounts[0] });

        const newCheckInTime = initialCheckInTime + 300; // Thêm 5 phút
        const newDetails = "Updated check-in";
        await tracker.updateAttendance(employeeId, newCheckInTime, newDetails, { from: accounts[0] });

        const updatedRecord = await tracker.getAttendanceByEmployee(employeeId);
        assert.equal(updatedRecord.checkInTime.toNumber(), newCheckInTime, "Check-in time was not updated");
        assert.equal(updatedRecord.details, newDetails, "Details were not updated");
    });
});