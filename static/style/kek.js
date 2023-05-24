
var Title = {
    name: "Title",
    value: null,
}
var SubTitle = {
    name: "SubTitle",
    value: null,
}
var AuthorName = {
    name: "AuthorName",
    value: null,
}
var AuthorPhoto = {
    name: "AuthorPhoto",
    value: null,
}
var Data = {
    name: "Data",
    value: null,
}
var HeroBigImage = {
    name: "HeroBigImage",
    value: null,
}
var HeroSmallImage = {
    name: "HeroSmallImage",
    value: null,
}
var Content = {
    name: "Content",
    value: null,
}

const readerAuthor = new FileReader();
const readerBig = new FileReader();
const readerSmall = new FileReader();

let
    userTitle = document.getElementById('title')
userSubTitle = document.getElementById('subtitle'),
    userAuthorName = document.getElementById('author-name'),
    userAuthorPhoto,
    authorIcon,
    userData,
    userBigImage,
    BImage,
    userSmallImage,
    SImage,
    userContent,

    getTitleArticle,
    getSubTitleArticle,
    getTitle,
    getSubTitle,
    getAuthorName,
    getAuthorPhoto,
    getData,
    getBigImage,
    getSmallImage;

function onload() {
    userTitle = document.getElementById('title');
    userSubTitle = document.getElementById('subtitle');
    userAuthorName = document.getElementById('author-name');
    authorIcon = document.getElementById('author-icon');
    userData = document.getElementById('data');
    BImage = document.getElementById('big-image');
    SImage = document.getElementById('small-image');
    userContent = document.getElementById('content');

    getTitleArticle = document.getElementById('preview-title-article');
    getTitle = document.getElementById('preview-title');
    getSubTitle = document.getElementById('preview-subtitle');
    getSubTitleArticle = document.getElementById('preview-subtitle-article');
    getAuthorName = document.getElementById('preview-author-name');
    getAuthorPhoto = document.getElementById('preview-author-photo');
    getData = document.getElementById('preview-data');
    getBigImage = document.getElementById('preview-big-image');
    getSmallImage = document.getElementById('preview-small-image');
}

function Click() {
    PrintToLog();
}

function PrintToLog() {
    if ((userTitle.value !== "") && (userSubTitle.value !== "") && (AuthorName.value !== "") && (authorIcon.src !== "file:///E:/web/Blog/static/svg_files/photo_icon.svg") && (userData.value !== "") && (BImage.src !== "file:///E:/web/Blog/static/images/hero_image_big.png") && (SImage.src !== "file:///E:/web/Blog/static/images/hero_image_small.png") && (Content !== "")) {
        Title.value = userTitle.value;
        console.log(Title.name, ':', Title.value);
        SubTitle.value = userSubTitle.value;
        console.log(SubTitle.name, ':', SubTitle.value);
        AuthorName.value = userAuthorName.value;
        console.log(AuthorName.name, ':', AuthorName.value);
        AuthorPhoto.value = authorIcon.src;
        console.log(AuthorPhoto.name, ':', AuthorPhoto.value);
        Data.value = userData.value;
        console.log(Data.name, ':', Data.value);
        HeroBigImage.value = BImage.src;
        console.log(HeroBigImage.name, ':', HeroBigImage.value);
        HeroSmallImage.value = SImage.src;
        console.log(HeroSmallImage.name, ':', HeroSmallImage.value);
        Content.value = userContent.value;
        console.log(Content.name, ':', Content.value);
        Preview();
    }
    else {
        console.log("Ошибка!!!! Некоторое поле пустое");
    }
}

function Preview() {
    getTitleArticle.innerHTML = Title.value;
    getTitle.innerHTML = Title.value;
    getSubTitleArticle.innerHTML = SubTitle.value;
    getSubTitle.innerHTML = SubTitle.value;
    getAuthorName.innerHTML = AuthorName.value;
    getAuthorPhoto.src = AuthorPhoto.value;
    getData.innerHTML = Data.value;
    getBigImage.src = HeroBigImage.value;
    getSmallImage.src = HeroSmallImage.value;
}

function ChangeIcon() {
    userAuthorPhoto = document.getElementById('author-photo').files[0];

    readerAuthor.addEventListener(
        "load",
        () => {
            authorIcon.src = readerAuthor.result;
        },
        false
    );

    if (userAuthorPhoto) {
        readerAuthor.readAsDataURL(userAuthorPhoto);
    }
}

function ChangeBigImage() {
    userBigImage = document.getElementById('hero-image-big').files[0];

    readerBig.addEventListener(
        "load",
        () => {
            BImage.src = readerBig.result;
        },
        false
    );

    if (userBigImage) {
        readerBig.readAsDataURL(userBigImage);
    }
}

function ChangeSmallImage() {
    userSmallImage = document.getElementById('hero-image-small').files[0];

    readerSmall.addEventListener(
        "load",
        () => {
            SImage.src = readerSmall.result;
        },
        false
    );

    if (userSmallImage) {
        readerSmall.readAsDataURL(userSmallImage);
    }
}



function uploadBig(image) {
    var reader = new FileReader();
    reader.addEventListener("load", function (e) {
        all_inputs.removeChild(big_blank);
        big_image.parentElement.insertAdjacentHTML(
            "beforeend",
            `<img class="big-image__size position__bs-image" src = ` + reader.result + `>`
        );
        first.insertAdjacentHTML(
            "beforeend",
            `<img src = ` + reader.result + `alt="" class="unload-image" >`
        );
    },
        false
    );
    reader.readAsDataURL(image);
}

function uploadSmall(image) {
    var reader = new FileReader();
    let kek;
    reader.addEventListener("load", function (e) {
        small_image.parentElement.insertAdjacentHTML(
            "beforeend",
            `<img class="small-image__size position__bs-image" src = ` + reader.result + `>`
        );
        second.insertAdjacentHTML(
            "beforeend",
            `<img src = ` + reader.result + `alt="" class="unload-image" >`
        );
    }, false);
    reader.readAsDataURL(image);
}




function upload_images(filer, filew) {
    var reader = new FileReader();
    reader.addEventListener("load", function (e) {
        filew.src = reader.result;
        big_blank.hidden = false;
        parent_big.insertAdjacentHTML(
            "beforeend",
            `<img class="small-image__size position__bs-image" src = ` + filew.src + `>`
        );

        first.insertAdjacentHTML(
            "beforeend",
            `<img src = ` + filew.src + `class="unload-image" >`
        );
    },
    );
    reader.readAsDataURL(filer);
}