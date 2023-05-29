function login() {
    var loginData = {
        email: document.getElementById("email").value,
        password: document.getElementById("pw").value
    };

    fetch('/login', {
        method: "POST",
        body: JSON.stringify(loginData),
        headers: { "content-type": "application/json; charset=UTF-8" }
    })
    .then(response => {
        if (response.ok) {
            window.location.href = "student.html"; // Redirect to student.html
        } else {
            throw new Error(response.statusText);
        }
    })
    .catch(error => {
        if (error.message === "Unauthorized") {
            alert("Credentials do not match!");
        } else {
            console.error("An error occurred:", error);
        }
    });
}

function logout() {
    fetch('/logout')
    .then(response => {
        if (response.ok) {
            window.open("index.html", "_self")
        } else {
            throw new Error(response.statusText)
        }
    }).catch (e => {
        alert(e)
    })
}