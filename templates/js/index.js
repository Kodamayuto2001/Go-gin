function main() {
    const fileInput = document.querySelector('input[type="file"]');
    const reader = new FileReader();
    const body = document.querySelector('body');
    const img = document.createElement('img');
    const button = document.createElement('button');
    button.innerHTML = 'upload';
    const formData = document.querySelector("form")
    const fd = new FormData(formData)

    function handleEvent(event) {
        img.src = reader.result;
        img.style.height = '300px';
        body.appendChild(img);
        body.appendChild(button);
    }

    function handleSelected(e) {
        const selectedFile = fileInput.files[0];
        console.log(selectedFile)
        if (selectedFile) {
            reader.addEventListener('load', handleEvent);
            reader.readAsDataURL(selectedFile);
        }
    }

    function requestListener(event){
        console.log("loadされました")
    }

    function request(event) {
        const selectedFile = fileInput.files[0];
        console.log(selectedFile)
        fd.append("image",selectedFile)
        
        let oRequest = new XMLHttpRequest();
        oRequest.open("POST","/upload")
        oRequest.send(fd)
        console.log("uploadされました")

        oRequest.addEventListener("load",requestListener)
    }

    fileInput.addEventListener('change', handleSelected);
    button.addEventListener('click', request);
}

main();
