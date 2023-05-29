function signUp() {
    var adminData = {
        firstname : document.getElementById("fname").value ,
        lastname : document.getElementById("lname").value,
        email : document.getElementById("email").value,
        password : document.getElementById("pw1").value,
        pw : document.getElementById("pw2").value,
    }
    if (adminData.password !== adminData.pw) {
        alert("PASSWORD doesn't match!")
        return
    }
    fetch('/signup', {
        method : "POST",
        body :JSON.stringify(adminData),
        headers: {"content-type" : "application/json; charset=UTF-8"}
    })
    .then((response) => {
        if (response.ok) {
          console.log("Redirecting to login page");
          window.open("index.html", "_self");
        } else {
          throw new Error("Network response was not ok.");
        }
      })
      .catch((error) => {
        console.error("Error:", error);
      });
}