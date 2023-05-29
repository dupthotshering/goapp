function resetform() {
  document.getElementById("cid").value = "";
  document.getElementById("cname").value = "";
}

window.onload = function() {
  fetch('/courses')
      .then(response => response.text())
      .then(data => showCourses(data));
}  

function getFormData() {
  var formData = {
      courseid : document.getElementById('cid').value,
      courseName : document.getElementById('cname').value,
  }
  return formData;
}

function addCourse() {
  var data = getFormData()

  if (data.stdid == "" || data.courseName == "") {
    alert("Course ID and name cannot be empty")
    return 
  }

  var cid  = data.courseid;

  fetch("/course" , {
      method : "POST",
      body: JSON.stringify(data),
      headers: {"Content-type": "application/json; charset=UTF-8"}

  }).then(response1 => {
      if (response1.ok) {
          fetch('/course/'+cid)
          .then(response2 => response2.text())
          .then(data => showCourse(data))
      } else {
          throw new Error(response1.statusText)
      }
  }).catch(e => {
      alert(e)
  })
  resetform();

}

function newRow(course) {
  var table = document.getElementById("myTable");
  var row = table.insertRow(table.length);

  var td = []
  for (i = 0; i < table.rows[0].cells.length; i++) {
      td[i] = row.insertCell(i);
  }
  td[0].innerHTML = course.courseid;
  td[1].innerHTML = course.coursename;
  td[2].innerHTML = '<input type="button" onclick="deleteCourse(this)" value="delete" id="button-1">';
  td[3].innerHTML = '<input type="button" onclick="updateCourse(this)" value="edit" id="button-2">';
}

function showCourse(data) {
  const course = JSON.parse(data);
  newRow(course)
}

function showCourses(data) {
  const courses = JSON.parse(data)
  courses.forEach(cour => {
      newRow(cour);
  });
}

var selectedRow = null


function updateCourse(r) {
  selectedRow = r.parentElement.parentElement;
  // console.log(selectedRow)

  document.getElementById("cid").value = selectedRow.cells[0].innerHTML;
  document.getElementById("cname").value = selectedRow.cells[1].innerHTML;

  var btn = document.getElementById("button-add");

  cid = selectedRow.cells[0].innerHTML;

  if (btn) {
      btn.innerHTML = "Update";
      btn.setAttribute("onclick" , "update(cid)");
  }
}

function update(cid) {
  var newData = getFormData()
  fetch('/course/'+cid,{
      method : "PUT",
      body : JSON.stringify(newData),
      headers : {"Content-type" : "application/json; charset=UTF-8"}

  }).then(res => {
      if(res.ok){
          selectedRow.cells[0].innerHTML = newData.courseid;
          selectedRow.cells[1].innerHTML = newData.courseName;

          var button = document.getElementById("button-add");
          button.innerHTML = "ADD";
          button.setAttribute("onclick", "addCourse()");

          selectedRow = null;
          resetform();

      }else {
          alert("Server : Update request error")
      }

  })
}

function deleteCourse(r) {
  if (confirm('Are you sure you want to DELETE this?')) {
      selectedRow = r.parentElement.parentElement;
      cid = selectedRow.cells[0].innerHTML;

      fetch('/course/'+cid,{
          method:"DELETE",
          headers : {"Content-type" : "application/json; charset=UTF-8"}
      });

      var rowIndex = selectedRow.rowIndex;
      if (rowIndex > 0) {
          document.getElementById("myTable").deleteRow(rowIndex);
      }
      selectedRow = null;

    }
}