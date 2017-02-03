{{define "layout"}}
<!DOCTYPE html>
<html lang="en">

<head>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Juan F. Giménez Silva - CV</title>

    <!-- Bootstrap Core CSS -->
    <link href="css/bootstrap.min.css" rel="stylesheet">
    <link href="css/main.css" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Assistant|Jaldi:700" rel="stylesheet"> 
    <!-- Custom CSS -->
    <style>
    body {
        padding-top: 15px;
        /* Required padding for .navbar-fixed-top. Remove if using .navbar-static-top. Change if height of navigation changes. */
    }
    </style>
   <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
        <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
        <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->

    <!-- Bootstrap Core JavaScript -->
    <script src="/js/jquery.js"></script>
    <script type="text/javascript" src="/js/bootstrap.min.js"></script>

    <script type="text/javascript" src="/js/react.js"></script>

    <script src="https://use.fontawesome.com/f88c962d3f.js"></script>

    <script type="text/javascript" src="/js/react-dom.js"></script>
    
    <script src="/js/browser.min.js"></script>

    <script src="/js/showdown.min.js"></script>
    
</head>

<body>

    <!-- Page Content -->
    <div class="container">
   {{template "body"}}
    </div>
    <!-- /.container -->



</body>

</html>
{{end}}
