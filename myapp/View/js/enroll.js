window.onload = function () {
    // request to get all students
    fetch('/students')
        .then(response => response.text())
        .then(data => getStudents(data));
    // get all courses
    fetch('/courses')
        .then(response => response.text())
        .then(data => getCourses(data));
    // get all enrolls
    fetch("/enrolls")
        .then(resp => resp.text())
        .then(data => getAllEnroll(data))
   
}

function getStudents(data) {
    const students = [] // to store all student ID
    const allStudents = JSON.parse(data)
    // add student id to students list
    allStudents.forEach(stud => {
        students.push(stud.stdid)
    });
    // get the sid input field
    var select = document.getElementById("sid")
    // iterate over the students and create a new option for each ID
    for (var i = 0; i < students.length; i++) {
        var sid = students[i];
        var option = document.createElement("option")
        option.textContent= sid;
        option.value = sid;
        select.appendChild(option);
    }
}

function getCourses(data) {
    const courses = [] 
    const allCourses = JSON.parse(data)
    allCourses.forEach(course => {
        courses.push(course.courseid)
    });
    var select = document.getElementById("cid")
    for (var i =0; i < courses.length; i++){
        var cid = courses[i]
        var option = document.createElement("option")
        option.innerHTML = cid
        option.value = cid
        select.appendChild(option)
    }
   
}

function addEnroll() {
    var _data = {
        stdid : parseInt(document.getElementById("sid").value),
        courseid : document.getElementById("cid").value,
        date: ""
    }
    var sid = _data.stdid;
    var cid = _data.courseid;

    if (isNaN(sid == "") || cid == "") {
        alert("Select valid data")
        return
    }
    fetch('/enroll', {
        method: "POST",
        body: JSON.stringify(_data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response => {
        if (response.ok) {
            fetch('/enroll/'+sid+'/'+cid)
            .then(response => response.text())
            .then(data => getEnrolled(data))
        } else {
            throw new Error(response.statusText)
        }
    }).catch(e => {
        if (e == "Error: Forbidden") {
            alert(e+". Duplicate entry!")
        }
    });
}

function getEnrolled(data) {
    const enrolled = JSON.parse(data)
    showTable(enrolled)
}

function showTable(enrolled) {
    var table = document.getElementById("myTable");
    var row = table.insertRow(table.length);
    var td=[]
    for(i=0; i<table.rows[0].cells.length; i++){
        td[i] = row.insertCell(i);
    }
    td[0].innerHTML = enrolled.stdid;
    td[1].innerHTML = enrolled.courseid;
    td[2].innerHTML = enrolled.date.split("T")[0];
    td[3].innerHTML = '<input type="button" onclick="deleteEnroll(this)" value="Delete" id="button-1">';
}

function getAllEnroll(data) {
    const allenroll = JSON.parse(data)
    console.log(allenroll)
    allenroll.forEach(enroll => {
        showTable(enroll)
    });
}

const deleteEnroll = async(r) => {
    if (confirm('Are you sure you want to DELETE this?')){
        selectedRow = r.parentElement.parentElement;
        sid = selectedRow.cells[0].innerHTML;
        cid = selectedRow.cells[1].innerHTML;
        fetch('/enroll/'+sid+"/"+cid, {
            method: "DELETE",
            headers: {"Content-type": "application/json; charset=UTF-8"}
        }). then( response => {
            if (response.ok) {
                var rowIndex = selectedRow.rowIndex;
                if (rowIndex>0) { 
                    document.getElementById("myTable").deleteRow(rowIndex);
                }
            }
        });
    }
}
