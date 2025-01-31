package main

const (
	compPage = `


<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="initial-scale=1, maximum-scale=1, user-scalable=no, width=device-width">
    <title>Exosphere Analytics Dashboard</title>
    <link href="/scripts/ionic.css" rel="stylesheet">
    <link href="/scripts/app.css" rel="stylesheet">
		<link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">
		<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.0.0-alpha1/jquery.min.js"></script>
			 <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>


</head>

<body>
  <nav class="navbar navbar-default">
  <div class="container-fluid">
    <!-- Brand and toggle get grouped for better mobile display -->
    <div class="navbar-header">
		<a class="navbar-brand" href="#">
		<img alt="Brand" src="http://www.vmware.com/files/images/vmrc/VMware_logo_gry_RGB_72dpi.jpg" height="25px">
		</a>
      <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
        <span class="sr-only">Toggle navigation</span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a class="navbar-brand" href="#">VMware Status Reporter</a>
    </div>

    <!-- Collect the nav links, forms, and other content for toggling -->
    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
		<ul class="nav navbar-nav">
			<li class="active"><a href="#">Help <span class="sr-only">(current)</span></a></li>
			<li><a href="#">About</a></li>
			<li class="dropdown">
				<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Tools <span class="caret"></span></a>
				<ul class="dropdown-menu">
					<li><a href="/admin">Admin</a></li>
					<li><a href="#">Help</a></li>
					<li><a href="/analytics">Analytics</a></li>
					<li><a href="/heatmap">Heatmap</a></li>
					<li role="separator" class="divider"></li>
					<li><a href="#">Send to Socialcast</a></li>
				</ul>
			</li>
			<li><a href="#"><span class="sr-only">(current)</span></a></li>
			<li><a href="#"><span class="sr-only">(current)</span></a></li>
			<li><a href="#"><span class="sr-only">(current)</span></a></li>
			<li><a href="#"><span class="sr-only">(current)</span></a></li>
			<li><a href="#"><span class="sr-only">(current)</span></a></li>
			<li><a href="#"><span class="sr-only">(current)</span></a></li>
			<li><a href="#"><b>September 30th, 2015</b>: Tim Callaghan</a></li>
			<li><img src="https://n2.cdn.socialcast.com/801245/socialcast.s3.amazonaws.com/tenants/5258/profile_photos/732278/tim_callaghan_square140.jpg?AWSAccessKeyId=AKIAISVYYXCGCXLJL2TQ&Expires=1445169600&Signature=TlZgwT7FtBEr2E1qCextWCpNMfc%3D" height=50px></li>
		</ul>

    </div><!-- /.navbar-collapse -->

  </div><!-- /.container-fluid -->

  </nav>
<div id="left-nav">
    <div class="welcome">Welcome, Timothy</div>
    <div class="title">Menu</div>
    <ul class="list">
        <li class="list-item">
            <a class="item-content" href="#dashboard/1">
                <i class="ion-stats-bars"></i>
                <h3>Competitive Dashboard</h3>
                <p>Competitive Trends</p>
            </a>
        </li>
        <li class="list-item">
            <a class="item-content" href="#dashboard/2">
                <i class="ion-pie-graph"></i>
                <h3>What's Not Working</h3>
                <p>Not Working Trends</p>
            </a>
        </li>
        <li class="list-item">
            <a class="item-content" href="#dashboard/3">
                <i class="ion-connection-bars"></i>
                <h3>What's Working</h3>
                <p>What's Good Trends</p>
            </a>
        </li>
    </ul>
</div>

<div id="container" class="transition">
    <header class="bar bar-header nav-bar bar-stable">
        <button id="btn-menu" class="button button-icon" style="font-size: 22px;"><i class="ion-navicon-round"></i></button>
        <h1 class="title">Project Exosphere Dashboard</h1>
    </header>
<div>
  hello there
</div>
    <div id="content"></div>
</div>

<script src="http://coenraets.org/apps/olympic-dashboard/lib/jquery.js"></script>
<script src="http://coenraets.org/apps/olympic-dashboard/lib/fastclick.js"></script>
<script src="http://coenraets.org/apps/olympic-dashboard/lib/chartjs/globalize.min.js"></script>
<script src="http://coenraets.org/apps/olympic-dashboard/lib/chartjs/dx.chartjs.js"></script>

<script src="/scripts/data.js"></script>
<script src="/scripts/data2.js"></script>
<script src="/scripts/dashboard1.js"></script>
<script src="/scripts/dashboard2.js"></script>
<script src="/scripts/dashboard3.js"></script>
<script src="/scripts/app.js"></script>

</body>
</html>
`
)
