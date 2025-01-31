package main

const (
	charts = `
  <!DOCTYPE html>
  <html>
    <head>
      <meta http-equiv="Content-type" content="text/html; charset=utf-8">
      <title>Testing Pie Chart</title>
      <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">
      <!--<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.0.0-alpha1/jquery.min.js"></script>-->
			<script src="https://code.jquery.com/jquery-2.1.4.min.js"></script>
			<script src="http://code.jquery.com/ui/1.11.4/jquery-ui.js"></script>
      <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
      <link href="http://cdnjs.cloudflare.com/ajax/libs/x-editable/1.5.0/bootstrap3-editable/css/bootstrap-editable.css" rel="stylesheet"/>
      <script src="http://cdnjs.cloudflare.com/ajax/libs/x-editable/1.5.0/bootstrap3-editable/js/bootstrap-editable.min.js"></script>

     <!--<script type="text/javascript" src="d3/d3.v2.js"></script>-->
      <script src="http://d3js.org/d3.v2.js"></script>
  	<!-- Note: I made good use of the sample code provided by the D3JS community and extended it to fit my needs to create this simple dashboard -->
      <style type="text/css">


  #pieChart {
  	position:absolute;
  	top:100px;
  	left:10px;
  	width:400px;
  	height: 400px;
  }

  #lineChart {
  	position:absolute;
  	top:110px;
  	left:410px;
  	height: 150px;
  }

  #barChart {
  	position:absolute;
  	top:260px;
  	left:410px;
  	height: 250px;
  }

  .slice {
     font-size: 12pt;
     font-family: Verdana;
     fill: white; //svg specific - instead of color
     font-weight: bold;
    		}

  /*for line chart*/
  .axis path, .axis line {
      fill: none;
      stroke: black;
      shape-rendering: crispEdges; //The shape-rendering property is an SVG attribute, used here to make sure our axis and its tick mark lines are pixel-perfect.
  }

  .line {
    fill: none;
    /*stroke: steelblue;*/
    stroke-width: 3px;
  }

  .dot {
    /*fill: white;*/
    /*stroke: steelblue;*/
    stroke-width: 1.5px;
    }


  .axis text {
      font-family: Verdana;
      font-size: 11px;
  }

  .title {
  	 font-family: Verdana;
      font-size: 15px;

  }

  .xAxis {
      font-family: verdana;
      font-size: 11px;
      fill: black;
  }

  .yAxis {
      font-family: verdana;
      font-size: 11px;
      fill: white;
  }


  table {
  	border-collapse:collapse;
  	border: 0px;
  	font-family: Verdana;
  	color: #5C5558;
  	font-size: 12px;
  	text-align: right;
  }

  td {
  	padding-left: 10px;
  }

  #lineChartTitle1 {
  	font-family: Verdana;
  	font-size  : 14px;
  	fill       : lightgrey;
  	font-weight: bold;
  	text-anchor: middle;
  }

  #lineChartTitle2 {
  	font-family: Verdana;
  	font-size  : 72px;
  	fill       : grey;
  	text-anchor: middle;
  	font-weight: bold;
  	/*font-style: italic;*/
  }

      </style>
    </head>
    <body>
      <nav class="navbar navbar-default">
      <div class="container-fluid">
        <!-- Brand and toggle get grouped for better mobile display -->
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="/jcham">VMware Weekly Reporter</a>
        </div>

        <!-- Collect the nav links, forms, and other content for toggling -->
        <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
          <ul class="nav navbar-nav">
            <li><a href="#">About</a></li>
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Tools <span class="caret"></span></a>
              <ul class="dropdown-menu">
                <li><a href="#">Management</a></li>
                <li><a href="#">Analytics</a></li>
                <li><a href="#">Integrations</a></li>
                <li role="separator" class="divider"></li>
                <li><a href="#">Settings link</a></li>
              </ul>
              <li class="active"><a href="#">Help</a></li>
            </li>
            <li><a href="#"><span class="sr-only">(current)</span></a></li>
            <li><a href="#"><span class="sr-only">(current)</span></a></li>
            <li><a href="#"><span class="sr-only">(current)</span></a></li>
            <li><a href="#"><span class="sr-only">(current)</span></a></li>
            <li><a href="#"><span class="sr-only">(current)</span></a></li>
            <li><a href="#"><span class="sr-only">(current)</span></a></li>
            <li><a href="#"><b>September 30th, 2015</b>: Jonathan Cham</a></li>
          </ul>

        </div><!-- /.navbar-collapse -->

      </div><!-- /.container-fluid -->

    </nav>
    <div class="container">

      <div id="pieChart"></div>
      <div id="barChart"></div>
      <div id="lineChart"></div>
      <script type="text/javascript">

  /*
  ################ FORMATS ##################
  -------------------------------------------
  */


  var 	formatAsPercentage = d3.format("%"),
  		formatAsPercentage1Dec = d3.format(".1%"),
  		formatAsInteger = d3.format(","),
  		fsec = d3.time.format("%S s"),
  		fmin = d3.time.format("%M m"),
  		fhou = d3.time.format("%H h"),
  		fwee = d3.time.format("%a"),
  		fdat = d3.time.format("%d d"),
  		fmon = d3.time.format("%b")
  		;

  /*
  ############# PIE CHART ###################
  -------------------------------------------
  */



  function dsPieChart(){

  	var dataset = [
  			{category: "Sam", measure: 0.30},
  	      {category: "Peter", measure: 0.25},
  	      {category: "John", measure: 0.15},
  	      {category: "Rick", measure: 0.05},
  	      {category: "Lenny", measure: 0.18},
  	      {category: "Paul", measure:0.04},
  	      {category: "Steve", measure: 0.03}
  	      ]
  	      ;

  	var 	width = 400,
  		   height = 400,
  		   outerRadius = Math.min(width, height) / 2,
  		   innerRadius = outerRadius * .999,
  		   // for animation
  		   innerRadiusFinal = outerRadius * .5,
  		   innerRadiusFinal3 = outerRadius* .45,
  		   color = d3.scale.category20()    //builtin range of colors
  		   ;

  	var vis = d3.select("#pieChart")
  	     .append("svg:svg")              //create the SVG element inside the <body>
  	     .data([dataset])                   //associate our data with the document
  	         .attr("width", width)           //set the width and height of our visualization (these will be attributes of the <svg> tag
  	         .attr("height", height)
  	     		.append("svg:g")                //make a group to hold our pie chart
  	         .attr("transform", "translate(" + outerRadius + "," + outerRadius + ")")    //move the center of the pie chart from 0, 0 to radius, radius
  				;

     var arc = d3.svg.arc()              //this will create <path> elements for us using arc data
          	.outerRadius(outerRadius).innerRadius(innerRadius);

     // for animation
     var arcFinal = d3.svg.arc().innerRadius(innerRadiusFinal).outerRadius(outerRadius);
  	var arcFinal3 = d3.svg.arc().innerRadius(innerRadiusFinal3).outerRadius(outerRadius);

     var pie = d3.layout.pie()           //this will create arc data for us given a list of values
          .value(function(d) { return d.measure; });    //we must tell it out to access the value of each element in our data array

     var arcs = vis.selectAll("g.slice")     //this selects all <g> elements with class slice (there aren't any yet)
          .data(pie)                          //associate the generated pie data (an array of arcs, each having startAngle, endAngle and value properties)
          .enter()                            //this will create <g> elements for every "extra" data element that should be associated with a selection. The result is creating a <g> for every object in the data array
              .append("svg:g")                //create a group to hold each slice (we will have a <path> and a <text> element associated with each slice)
                 .attr("class", "slice")    //allow us to style things in the slices (like text)
                 .on("mouseover", mouseover)
      				.on("mouseout", mouseout)
      				.on("click", up)
      				;

          arcs.append("svg:path")
                 .attr("fill", function(d, i) { return color(i); } ) //set the color for each slice to be chosen from the color function defined above
                 .attr("d", arc)     //this creates the actual SVG path using the associated data (pie) with the arc drawing function
  					.append("svg:title") //mouseover title showing the figures
  				   .text(function(d) { return d.data.category + ": " + formatAsPercentage(d.data.measure); });

          d3.selectAll("g.slice").selectAll("path").transition()
  			    .duration(750)
  			    .delay(10)
  			    .attr("d", arcFinal )
  			    ;

  	  // Add a label to the larger arcs, translated to the arc centroid and rotated.
  	  // source: http://bl.ocks.org/1305337#index.html
  	  arcs.filter(function(d) { return d.endAngle - d.startAngle > .2; })
  	  		.append("svg:text")
  	      .attr("dy", ".35em")
  	      .attr("text-anchor", "middle")
  	      .attr("transform", function(d) { return "translate(" + arcFinal.centroid(d) + ")rotate(" + angle(d) + ")"; })
  	      //.text(function(d) { return formatAsPercentage(d.value); })
  	      .text(function(d) { return d.data.category; })
  	      ;

  	   // Computes the label angle of an arc, converting from radians to degrees.
  		function angle(d) {
  		    var a = (d.startAngle + d.endAngle) * 90 / Math.PI - 90;
  		    return a > 90 ? a - 180 : a;
  		}


  		// Pie chart title
  		vis.append("svg:text")
  	     	.attr("dy", ".35em")
  	      .attr("text-anchor", "middle")
  	      .text("Revenue Share 2012")
  	      .attr("class","title")
  	      ;



  	function mouseover() {
  	  d3.select(this).select("path").transition()
  	      .duration(750)
  	        		//.attr("stroke","red")
  	        		//.attr("stroke-width", 1.5)
  	        		.attr("d", arcFinal3)
  	        		;
  	}

  	function mouseout() {
  	  d3.select(this).select("path").transition()
  	      .duration(750)
  	        		//.attr("stroke","blue")
  	        		//.attr("stroke-width", 1.5)
  	        		.attr("d", arcFinal)
  	        		;
  	}

  	function up(d, i) {

  				/* update bar chart when user selects piece of the pie chart */
  				//updateBarChart(dataset[i].category);
  				updateBarChart(d.data.category, color(i));
  				updateLineChart(d.data.category, color(i));

  	}
  }

  dsPieChart();

  /*
  ############# BAR CHART ###################
  -------------------------------------------
  */



  var datasetBarChart = [
  //{ group: "All", category: "Oranges", measure: 63850.4963 },
  //{ group: "All", category: "Apples", measure: 78258.0845 },
  { group: "All", category: "Grapes", measure: 60610.2355 },
  { group: "All", category: "Figs", measure: 30493.1686 },
  { group: "All", category: "Mangos", measure: 56097.0151 },
  { group: "Sam", category: "Oranges", measure: 19441.5648 },
  { group: "Sam", category: "Apples", measure: 25922.0864 },
  { group: "Sam", category: "Grapes", measure: 9720.7824 },
  { group: "Sam", category: "Figs", measure: 6480.5216 },
  { group: "Sam", category: "Mangos", measure: 19441.5648 },
  { group: "Peter", category: "Oranges", measure: 22913.2728 },
  { group: "Peter", category: "Apples", measure: 7637.7576 },
  { group: "Peter", category: "Grapes", measure: 23549.7526 },
  { group: "Peter", category: "Figs", measure: 1909.4394 },
  { group: "Peter", category: "Mangos", measure: 7637.7576 },
  { group: "John", category: "Oranges", measure: 1041.5124 },
  { group: "John", category: "Apples", measure: 2430.1956 },
  { group: "John", category: "Grapes", measure: 15275.5152 },
  { group: "John", category: "Figs", measure: 4166.0496 },
  { group: "John", category: "Mangos", measure: 11803.8072 },
  { group: "Rick", category: "Oranges", measure: 7406.3104 },
  { group: "Rick", category: "Apples", measure: 2545.9192 },
  { group: "Rick", category: "Grapes", measure: 1620.1304 },
  { group: "Rick", category: "Figs", measure: 8563.5464 },
  { group: "Rick", category: "Mangos", measure: 3008.8136 },
  { group: "Lenny", category: "Oranges", measure: 7637.7576 },
  { group: "Lenny", category: "Apples", measure: 35411.4216 },
  { group: "Lenny", category: "Grapes", measure: 8332.0992 },
  { group: "Lenny", category: "Figs", measure: 6249.0744 },
  { group: "Lenny", category: "Mangos", measure: 11803.8072 },
  { group: "Paul", category: "Oranges", measure: 3182.399 },
  { group: "Paul", category: "Apples", measure: 867.927 },
  { group: "Paul", category: "Grapes", measure: 1808.18125 },
  { group: "Paul", category: "Figs", measure: 795.59975 },
  { group: "Paul", category: "Mangos", measure: 578.618 },
  { group: "Steve", category: "Oranges", measure: 2227.6793 },
  { group: "Steve", category: "Apples", measure: 3442.7771 },
  { group: "Steve", category: "Grapes", measure: 303.77445 },
  { group: "Steve", category: "Figs", measure: 2328.93745 },
  { group: "Steve", category: "Mangos", measure: 1822.6467 },
  ]
  ;

  // set initial group value
  var group = "All";

  function datasetBarChosen(group) {
  	var ds = [];
  	for (x in datasetBarChart) {
  		 if(datasetBarChart[x].group==group){
  		 	ds.push(datasetBarChart[x]);
  		 }
  		}
  	return ds;
  }


  function dsBarChartBasics() {

  		var margin = {top: 30, right: 5, bottom: 20, left: 50},
  		width = 500 - margin.left - margin.right,
  	   height = 250 - margin.top - margin.bottom,
  		colorBar = d3.scale.category20(),
  		barPadding = 1
  		;

  		return {
  			margin : margin,
  			width : width,
  			height : height,
  			colorBar : colorBar,
  			barPadding : barPadding
  		}
  		;
  }

  function dsBarChart() {

  	var firstDatasetBarChart = datasetBarChosen(group);

  	var basics = dsBarChartBasics();

  	var margin = basics.margin,
  		width = basics.width,
  	   height = basics.height,
  		colorBar = basics.colorBar,
  		barPadding = basics.barPadding
  		;

  	var 	xScale = d3.scale.linear()
  						.domain([0, firstDatasetBarChart.length])
  						.range([0, width])
  						;

  	// Create linear y scale
  	// Purpose: No matter what the data is, the bar should fit into the svg area; bars should not
  	// get higher than the svg height. Hence incoming data needs to be scaled to fit into the svg area.
  	var yScale = d3.scale.linear()
  			// use the max funtion to derive end point of the domain (max value of the dataset)
  			// do not use the min value of the dataset as min of the domain as otherwise you will not see the first bar
  		   .domain([0, d3.max(firstDatasetBarChart, function(d) { return d.measure; })])
  		   // As coordinates are always defined from the top left corner, the y position of the bar
  		   // is the svg height minus the data value. So you basically draw the bar starting from the top.
  		   // To have the y position calculated by the range function
  		   .range([height, 0])
  		   ;

  	//Create SVG element

  	var svg = d3.select("#barChart")
  			.append("svg")
  		    .attr("width", width + margin.left + margin.right)
  		    .attr("height", height + margin.top + margin.bottom)
  		    .attr("id","barChartPlot")
  		    ;

  	var plot = svg
  		    .append("g")
  		    .attr("transform", "translate(" + margin.left + "," + margin.top + ")")
  		    ;

  	plot.selectAll("rect")
  		   .data(firstDatasetBarChart)
  		   .enter()
  		   .append("rect")
  			.attr("x", function(d, i) {
  			    return xScale(i);
  			})
  		   .attr("width", width / firstDatasetBarChart.length - barPadding)
  			.attr("y", function(d) {
  			    return yScale(d.measure);
  			})
  			.attr("height", function(d) {
  			    return height-yScale(d.measure);
  			})
  			.attr("fill", "lightgrey")
  			;


  	// Add y labels to plot

  	plot.selectAll("text")
  	.data(firstDatasetBarChart)
  	.enter()
  	.append("text")
  	.text(function(d) {
  			return formatAsInteger(d3.round(d.measure));
  	})
  	.attr("text-anchor", "middle")
  	// Set x position to the left edge of each bar plus half the bar width
  	.attr("x", function(d, i) {
  			return (i * (width / firstDatasetBarChart.length)) + ((width / firstDatasetBarChart.length - barPadding) / 2);
  	})
  	.attr("y", function(d) {
  			return yScale(d.measure) + 14;
  	})
  	.attr("class", "yAxis")
  	/* moved to CSS
  	.attr("font-family", "sans-serif")
  	.attr("font-size", "11px")
  	.attr("fill", "white")
  	*/
  	;

  	// Add x labels to chart

  	var xLabels = svg
  		    .append("g")
  		    .attr("transform", "translate(" + margin.left + "," + (margin.top + height)  + ")")
  		    ;

  	xLabels.selectAll("text.xAxis")
  		  .data(firstDatasetBarChart)
  		  .enter()
  		  .append("text")
  		  .text(function(d) { return d.category;})
  		  .attr("text-anchor", "middle")
  			// Set x position to the left edge of each bar plus half the bar width
  						   .attr("x", function(d, i) {
  						   		return (i * (width / firstDatasetBarChart.length)) + ((width / firstDatasetBarChart.length - barPadding) / 2);
  						   })
  		  .attr("y", 15)
  		  .attr("class", "xAxis")
  		  //.attr("style", "font-size: 12; font-family: Helvetica, sans-serif")
  		  ;

  	// Title

  	svg.append("text")
  		.attr("x", (width + margin.left + margin.right)/2)
  		.attr("y", 15)
  		.attr("class","title")
  		.attr("text-anchor", "middle")
  		.text("Overall Sales Breakdown 2012")
  		;
  }

  dsBarChart();

   /* ** UPDATE CHART ** */

  /* updates bar chart on request */

  function updateBarChart(group, colorChosen) {

  		var currentDatasetBarChart = datasetBarChosen(group);

  		var basics = dsBarChartBasics();

  		var margin = basics.margin,
  			width = basics.width,
  		   height = basics.height,
  			colorBar = basics.colorBar,
  			barPadding = basics.barPadding
  			;

  		var 	xScale = d3.scale.linear()
  			.domain([0, currentDatasetBarChart.length])
  			.range([0, width])
  			;


  		var yScale = d3.scale.linear()
  	      .domain([0, d3.max(currentDatasetBarChart, function(d) { return d.measure; })])
  	      .range([height,0])
  	      ;

  	   var svg = d3.select("#barChart svg");

  	   var plot = d3.select("#barChartPlot")
  	   	.datum(currentDatasetBarChart)
  		   ;

  	  		/* Note that here we only have to select the elements - no more appending! */
  	  	plot.selectAll("rect")
  	      .data(currentDatasetBarChart)
  	      .transition()
  			.duration(750)
  			.attr("x", function(d, i) {
  			    return xScale(i);
  			})
  		   .attr("width", width / currentDatasetBarChart.length - barPadding)
  			.attr("y", function(d) {
  			    return yScale(d.measure);
  			})
  			.attr("height", function(d) {
  			    return height-yScale(d.measure);
  			})
  			.attr("fill", colorChosen)
  			;

  		plot.selectAll("text.yAxis") // target the text element(s) which has a yAxis class defined
  			.data(currentDatasetBarChart)
  			.transition()
  			.duration(750)
  		   .attr("text-anchor", "middle")
  		   .attr("x", function(d, i) {
  		   		return (i * (width / currentDatasetBarChart.length)) + ((width / currentDatasetBarChart.length - barPadding) / 2);
  		   })
  		   .attr("y", function(d) {
  		   		return yScale(d.measure) + 14;
  		   })
  		   .text(function(d) {
  				return formatAsInteger(d3.round(d.measure));
  		   })
  		   .attr("class", "yAxis")
  		;


  		svg.selectAll("text.title") // target the text element(s) which has a title class defined
  			.attr("x", (width + margin.left + margin.right)/2)
  			.attr("y", 15)
  			.attr("class","title")
  			.attr("text-anchor", "middle")
  			.text(group + "'s Sales Breakdown 2012")
  		;
  }


  /*
  ############# LINE CHART ##################
  -------------------------------------------
  */

  var datasetLineChart = [
  { group: "All", category: 2008, measure: 289309 },
  { group: "All", category: 2009, measure: 234998 },
  { group: "All", category: 2010, measure: 310900 },
  { group: "All", category: 2011, measure: 223900 },
  { group: "All", category: 2012, measure: 234500 },
  { group: "Sam", category: 2008, measure: 81006.52 },
  { group: "Sam", category: 2009, measure: 70499.4 },
  { group: "Sam", category: 2010, measure: 96379 },
  { group: "Sam", category: 2011, measure: 64931 },
  { group: "Sam", category: 2012, measure: 70350 },
  { group: "Peter", category: 2008, measure: 63647.98 },
  { group: "Peter", category: 2009, measure: 61099.48 },
  { group: "Peter", category: 2010, measure: 87052 },
  { group: "Peter", category: 2011, measure: 58214 },
  { group: "Peter", category: 2012, measure: 58625 },
  { group: "Rick", category: 2008, measure: 23144.72 },
  { group: "Rick", category: 2009, measure: 14099.88 },
  { group: "Rick", category: 2010, measure: 15545 },
  { group: "Rick", category: 2011, measure: 11195 },
  { group: "Rick", category: 2012, measure: 11725 },
  { group: "John", category: 2008, measure: 34717.08 },
  { group: "John", category: 2009, measure: 30549.74 },
  { group: "John", category: 2010, measure: 34199 },
  { group: "John", category: 2011, measure: 33585 },
  { group: "John", category: 2012, measure: 35175 },
  { group: "Lenny", category: 2008, measure: 69434.16 },
  { group: "Lenny", category: 2009, measure: 46999.6 },
  { group: "Lenny", category: 2010, measure: 62180 },
  { group: "Lenny", category: 2011, measure: 40302 },
  { group: "Lenny", category: 2012, measure: 42210 },
  { group: "Paul", category: 2008, measure: 7232.725 },
  { group: "Paul", category: 2009, measure: 4699.96 },
  { group: "Paul", category: 2010, measure: 6218 },
  { group: "Paul", category: 2011, measure: 8956 },
  { group: "Paul", category: 2012, measure: 9380 },
  { group: "Steve", category: 2008, measure: 10125.815 },
  { group: "Steve", category: 2009, measure: 7049.94 },
  { group: "Steve", category: 2010, measure: 9327 },
  { group: "Steve", category: 2011, measure: 6717 },
  { group: "Steve", category: 2012, measure: 7035 }
  ]
  ;

  // set initial category value
  var group = "All";

  function datasetLineChartChosen(group) {
  	var ds = [];
  	for (x in datasetLineChart) {
  		 if(datasetLineChart[x].group==group){
  		 	ds.push(datasetLineChart[x]);
  		 }
  		}
  	return ds;
  }

  function dsLineChartBasics() {

  	var margin = {top: 20, right: 10, bottom: 0, left: 50},
  	    width = 500 - margin.left - margin.right,
  	    height = 150 - margin.top - margin.bottom
  	    ;

  		return {
  			margin : margin,
  			width : width,
  			height : height
  		}
  		;
  }


  function dsLineChart() {

  	var firstDatasetLineChart = datasetLineChartChosen(group);

  	var basics = dsLineChartBasics();

  	var margin = basics.margin,
  		width = basics.width,
  	   height = basics.height
  		;

  	var xScale = d3.scale.linear()
  	    .domain([0, firstDatasetLineChart.length-1])
  	    .range([0, width])
  	    ;

  	var yScale = d3.scale.linear()
  	    .domain([0, d3.max(firstDatasetLineChart, function(d) { return d.measure; })])
  	    .range([height, 0])
  	    ;

  	var line = d3.svg.line()
  	    //.x(function(d) { return xScale(d.category); })
  	    .x(function(d, i) { return xScale(i); })
  	    .y(function(d) { return yScale(d.measure); })
  	    ;

  	var svg = d3.select("#lineChart").append("svg")
  	    .datum(firstDatasetLineChart)
  	    .attr("width", width + margin.left + margin.right)
  	    .attr("height", height + margin.top + margin.bottom)
  	    // create group and move it so that margins are respected (space for axis and title)

  	var plot = svg
  	    .append("g")
  	    .attr("transform", "translate(" + margin.left + "," + margin.top + ")")
  	    .attr("id", "lineChartPlot")
  	    ;

  		/* descriptive titles as part of plot -- start */
  	var dsLength=firstDatasetLineChart.length;

  	plot.append("text")
  		.text(firstDatasetLineChart[dsLength-1].measure)
  		.attr("id","lineChartTitle2")
  		.attr("x",width/2)
  		.attr("y",height/2)
  		;
  	/* descriptive titles -- end */

  	plot.append("path")
  	    .attr("class", "line")
  	    .attr("d", line)
  	    // add color
  		.attr("stroke", "lightgrey")
  	    ;

  	plot.selectAll(".dot")
  	    .data(firstDatasetLineChart)
  	  	 .enter().append("circle")
  	    .attr("class", "dot")
  	    //.attr("stroke", function (d) { return d.measure==datasetMeasureMin ? "red" : (d.measure==datasetMeasureMax ? "green" : "steelblue") } )
  	    .attr("fill", function (d) { return d.measure==d3.min(firstDatasetLineChart, function(d) { return d.measure; }) ? "red" : (d.measure==d3.max(firstDatasetLineChart, function(d) { return d.measure; }) ? "green" : "white") } )
  	    //.attr("stroke-width", function (d) { return d.measure==datasetMeasureMin || d.measure==datasetMeasureMax ? "3px" : "1.5px"} )
  	    .attr("cx", line.x())
  	    .attr("cy", line.y())
  	    .attr("r", 3.5)
  	    .attr("stroke", "lightgrey")
  	    .append("title")
  	    .text(function(d) { return d.category + ": " + formatAsInteger(d.measure); })
  	    ;

  	svg.append("text")
  		.text("Performance 2012")
  		.attr("id","lineChartTitle1")
  		.attr("x",margin.left + ((width + margin.right)/2))
  		.attr("y", 10)
  		;

  }

  dsLineChart();


   /* ** UPDATE CHART ** */

  /* updates bar chart on request */
  function updateLineChart(group, colorChosen) {

  	var currentDatasetLineChart = datasetLineChartChosen(group);

  	var basics = dsLineChartBasics();

  	var margin = basics.margin,
  		width = basics.width,
  	   height = basics.height
  		;

  	var xScale = d3.scale.linear()
  	    .domain([0, currentDatasetLineChart.length-1])
  	    .range([0, width])
  	    ;

  	var yScale = d3.scale.linear()
  	    .domain([0, d3.max(currentDatasetLineChart, function(d) { return d.measure; })])
  	    .range([height, 0])
  	    ;

  	var line = d3.svg.line()
      .x(function(d, i) { return xScale(i); })
      .y(function(d) { return yScale(d.measure); })
      ;

     var plot = d3.select("#lineChartPlot")
     	.datum(currentDatasetLineChart)
  	   ;

  	/* descriptive titles as part of plot -- start */
  	var dsLength=currentDatasetLineChart.length;

  	plot.select("text")
  		.text(currentDatasetLineChart[dsLength-1].measure)
  		;
  	/* descriptive titles -- end */

  	plot
  	.select("path")
  		.transition()
  		.duration(750)
  	   .attr("class", "line")
  	   .attr("d", line)
  	   // add color
  		.attr("stroke", colorChosen)
  	   ;

  	var path = plot
  		.selectAll(".dot")
  	   .data(currentDatasetLineChart)
  	   .transition()
  		.duration(750)
  	   .attr("class", "dot")
  	   .attr("fill", function (d) { return d.measure==d3.min(currentDatasetLineChart, function(d) { return d.measure; }) ? "red" : (d.measure==d3.max(currentDatasetLineChart, function(d) { return d.measure; }) ? "green" : "white") } )
  	   .attr("cx", line.x())
  	   .attr("cy", line.y())
  	   .attr("r", 3.5)
  	   // add color
  		.attr("stroke", colorChosen)
  	   ;

  	   path
  	   .selectAll("title")
  	   .text(function(d) { return d.category + ": " + formatAsInteger(d.measure); })
  	   ;

  }

      </script>
    </div>
    </body>
  </html>

`
)
