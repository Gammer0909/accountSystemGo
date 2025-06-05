document.addEventListener("DOMContentLoaded", () => {
    console.log("DOM fully loaded");

    const formElem = document.getElementById("signup");

    if (!formElem) {
        console.error("Form element not found");
        return;
    }

    formElem.addEventListener("submit", async (e) => {
        e.preventDefault();  // Prevent form from submitting traditionally

        // Collect input values immediately after preventing default behavior
        const username = document.getElementById("username")?.value;
        const email = document.getElementById("email")?.value;
        const password = document.getElementById("password")?.value;

        // Defensive check to avoid null
        if (!username || !email || !password) {
            console.error("One or more input fields are missing");
            return;
        }

        try {
            const res = await fetch("http://localhost:8080/api/signup", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(
                    {
                        username,
                        email,
                        password
                    }
                )
            });
            console.log(res.status)

            if (res.status == 200) {
                // Success: parse response and redirect
                window.location.href = "../dashboard/dashboard.html";
            } else if (res.status == 409) {
                const text = await res.text();
                alert(text);
            }
        } catch (err) {
            console.error("Network Error:", err);
        }
    });
});
