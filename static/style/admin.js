//document.addEventListener('DOMContentLoaded', function () {
const form = document.getElementById('form');
form.addEventListener('sumbit', formSend);

async function formSend(e) {
    let error = formValidate(form);

    if (error === 0) {
        //в сервак
        Preview();
        alert('корректный ввод данных');
    } else {
        alert("не все данные верны, перепроверьте пожалуйста")
    }
}

function formValidate(form) {
    let error = 0;
    let formReq = document.querySelectorAll('._req');
    let formReq1 = document.querySelectorAll('._req1')
    for (let index = 0; index < formReq.length; index++) {
        const input = formReq[index];
        formRemoveError(input)
        if (input.value === "") {
            formAddError(input);
            error++;
        }
    };

    for (let index = 0; index < formReq1.length; index++) {
        const input = formReq1[index];
        formRemoveErrorFile(input)
        if (input.value === "") {
            formAddErrorFile(input);
            error++;
        }
    };
    return error;
}

function formAddError(input) {
    input.classList.add('_error');
}

function formRemoveError(input) {
    input.classList.remove('_error');
}

function formAddErrorFile(input) {
    input.parentElement.classList.add('_error');
}

function formRemoveErrorFile(input) {
    input.parentElement.classList.remove('_error');
}

//take image
const author_url = document.getElementById("author-image");
const author_preview = document.getElementById("previewA");
const big_image = document.getElementById("big-image");
const small_image = document.getElementById("small-image");
const parent_big = document.getElementsByClassName("input__big-image");
const first = document.getElementById("preview-post-one");
const second = document.getElementById("preview-post-two");
//take text
const title = document.getElementsByName("titleInput");
const subtitle = document.getElementsByName("subtitleInput");
const date = document.getElementsByName("dateInput");
const author_name = document.getElementsByName("AnameInput");
let
    preview_title = document.getElementsByName("preview_title"),
    preview_subtitle = document.getElementsByName("preview_subtitle"),
    preview_date = document.getElementsByName("preview_date"),
    preview_author_name = document.getElementsByName("preview_author_name");
//take blank
const block1 = document.getElementById("block1");
const block2 = document.getElementById("block2");
const block3 = document.getElementById("change_au-image");
/*content_remove.addEventListener("click",
    function (e) {
        content_remove.parentElement.insertAdjacentHTML(
            `<div class="big-image__size position__bs-image" name="Bblank">
            <img src={{.ImageCamera}} alt="">
            <p class="text-upload">{{.Upload}}</p>
        </div>`
        );
        content_remove.parentElement.children[1].remove();

    });*/


function Preview() {
    preview_title[0].textContent = title[0].value;
    preview_title[1].textContent = title[0].value;
    preview_author_name[0].textContent = author_name[0].value;
    preview_date[0].textContent = date[0].value;
    preview_subtitle[0].textContent = subtitle[0].value;
    preview_subtitle[1].textContent = subtitle[0].value;
};

author_url.addEventListener("change", () => {
    uploadFile(author_url.files[0], author_preview);
});


big_image.addEventListener("change", () => {
    uploadBig(big_image.files[0], first);
    block1.children[0].remove();
});

small_image.addEventListener("change", () => {
    uploadSmall(small_image.files[0], second);
    block2.children[0].remove();
});

function uploadFile(file, filewrite) {
    var reader = new FileReader();
    reader.addEventListener("load", () => {
        filewrite.src = reader.result;
        if (reader.result != "") {
            block3.src = reader.result;
        };
    },
        false
    );
    reader.readAsDataURL(file);

};


function uploadBig(image) {
    var reader = new FileReader();
    reader.addEventListener("load", function (e) {
        block1.insertAdjacentHTML(
            "beforeend",
            `
            <div>
                <img class="big-image__size position__bs-image" src = ` + reader.result + `>
                <div class="position-for-url andpadding">
                    <div>
                        <img class="" src="../static/image/camera.png">
                        <span>Upload New </span>
                    </div>
                    <div id="trash-remove">
                        <img src="../static/image/trash-2.png">
                        <span>Remove </span>
                    </div>
                </div>
            </div>
            `
        );
        first.src = reader.result;
    },
        false
    );
    reader.readAsDataURL(image);
};

function uploadSmall(image) {
    var reader = new FileReader();
    reader.addEventListener("load", function (e) {
        block2.insertAdjacentHTML(
            "beforeend",
            `
            <div>
                <img class="small-image__size position__bs-image" src = ` + reader.result + `>
                <div class="position-for-url andpadding">
                    <div>
                        <img class="" src="../static/image/camera.png">
                        <span>Upload New </span>
                    </div>
                    <div id="trash-remove">
                        <img id="trash-remove" src="../static/image/trash-2.png">
                        <span>Remove </span>
                    </div>
                </div>
            </div>
            `
        );
        second.src = reader.result;
    },
        false
    );
    reader.readAsDataURL(image);
}

//});