<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Kampus Biru</title>

    <link rel="stylesheet" href="bulma/css/bulma.min.css">
    <!-- <script defer src="https://use.fontawesome.com/releases/v5.3.1/js/all.js"></script> -->
</head>
<body>

    <?php include "header.html"; ?>

    <section class="hero">
        <div class="hero-body" style="padding:20px">
            <div class="box" box-padding="10px">
            <?php 
                if(@$_GET['page'] == "") {
                    include "depan.html";
                } else {
                    $page = $_GET['page'];

                    if($page == "tambahJurusan") {
                        include "tambahJurusan.html";
                    } elseif($page == "dataJurusan") {
                        include "dataJurusan.html";
                    }
                }
            ?>
            </div>
            
        </div>
    </section>

    <?php include "footer.html" ?>
</body>
</html>