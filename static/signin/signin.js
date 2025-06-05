document.addEventListener("DOMContentLoaded", () => {
    console.log("DOM fully loaded");

    const status = document.getElementById("status")
    const formElem = document.getElementById("login");

    if (!formElem) {    
        console.error("Form element not found");
        return;
    }

    formElem.addEventListener("submit", async (e) => {
        e.preventDefault();  // Prevent form from submitting traditionally

        // Collect input values immediately after preventing default behavior
        const email = document.getElementById("email")?.value;
        const password = document.getElementById("password")?.value;

        // Defensive check to avoid null
        if (!email || !password) {
            console.error("One or more input fields are missing");
            return;
        }

        const filler = ""

        try {
            const res = await fetch("http://localhost:8080/api/signin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(
                    {
                        filler,
                        email,
                        password
                    }
                )
            });

            console.log(res.status)

            if (res.status === 200) {
                // Success: parse response and redirect
                window.location.href = "../dashboard/dashboard.html";
            } else if (res.status === 401) {
                alert("Login failed: Check your username or password.");
            } else {
                // Other errors
                const text = await res.text();
                alert("Unexpected error (" + res.status + "): " + text);
            }
        } catch (err) {
            status.textContent = err.body
            console.log(err)
        }
    });
});
