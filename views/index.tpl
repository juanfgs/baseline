{{define "body"}}
<div id="container">
    <div class="row panel">
	<div class="col-md-4 bg_blur ">
	</div>
        <div class="col-md-8  col-xs-12">
            <img src="/images/author.png" class="img-thumbnail picture hidden-xs" />
            <img src="/images/author.png" class="img-thumbnail visible-xs picture_mob" />
            <div class="header">
                <h1>Juan F. Giménez Silva</h1>
                <h4>Web Developer / PHP, JS,Golang, Ruby developer / Free software enthusiast </h4>
                <span>My name is Juan, I'm from Argentina and I'm $age years old. I've been playing, learning and working with computers since I was 4 years old. Formally I've been working as a web developer for 11 years, using mostly PHP and JS. I love programming and software development in general. I use, create and encourage the use of Free Software whenver possible. </span>
            </div>
        </div>
    </div>
    <div class="row div_margin">
            <script type="text/babel" src="js/main.js"></script>
            <div id="cv-content"></div>
    </div>
</div>
{{end}}
