const init = () => {
    const $elSeedCount = document.getElementById("el-input-count");
    const $elSeedButton = document.getElementById("el-btn-seed");
    const $elSeedFeedback = document.getElementById("el-seed-feedback");

    const clearFeedback = () => {
        if ($elSeedFeedback.firstChild) {
            $elSeedFeedback.firstChild.remove();
        }
    };

    const setErrorFeedback = (msg) => {
        const alert = document.createElement("div");
        alert.classList.add("alert", "alert-danger");
        alert.innerText = `Error: ${msg}`;
        $elSeedFeedback.appendChild(alert);
    };

    const setSuccessFeedback = () => {
        const alert = document.createElement("div");
        alert.classList.add("alert", "alert-success");
        alert.innerText = `Operation was successful`;
        $elSeedFeedback.appendChild(alert);
    };

    $elSeedButton.addEventListener("click", async e => {
        e.preventDefault();
        clearFeedback();

        try {
            const res = await fetch("/api/seed", {
                method: "POST",
                headers: {
                    "content-type": "application/json",
                },
                body: JSON.stringify({
                    count: +$elSeedCount.value,
                })
            })

            if (!res.ok) {
                const message = await res.json();
                setErrorFeedback(message.error);
                return;
            }

            setSuccessFeedback();
        } catch (err) {
            setErrorFeedback(`unable to communicate with server: ${err}`);
        }
    });
}