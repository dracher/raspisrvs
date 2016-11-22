/*
Use your favourite build tool to fill this file which will contain the
../app/*.js and the .../vendor/*.js contents. 

I suggest you do it using 'gulp', but its your decision, you can use webpack also.
*/

$(function () {

  $('.ui.accordion')
    .accordion()
    ;

  var pathname = window.location.pathname;

  if (pathname === "/aqi") {
    getAQIData()
  };

  if (pathname === "/pistatus") {
    w = new Ws("ws://" + HOST + "/ws");

    w.OnConnect(function () {
      console.log("Websocket connection established");
    });

    w.On("cpu", function (message) {
      // console.log(message)
      $("#cpu_t").text(message)
    });
  };
});

var getAQIData = function () {
  var aqi_url = "/api/v1/airindex/";
  var city_list = [
    "beijing",
    "chengdu",
    "guangzhou",
    "shanghai",
    "shenyang",
  ];
  var city_name_mapper = {
    "beijing": "北京",
    "chengdu": "成都",
    "guangzhou": "广州",
    "shanghai": "上海",
    "shenyang": "沈阳"
  };

  city_list.forEach(function (value, index) {
    $.getJSON(aqi_url + value, function (data) {
      ret = data.resp

      var t = ret[0].reverse()
      var p = ret[2].map(parseFloat).reverse()
      var c = ret[1].map(parseFloat).reverse()

      var myChart = Highcharts.chart(value, {

        credits: {
            enabled: false
        },

        title: {
          text: ret[0][23] + ' ' + city_name_mapper[value] + ' AQI is  ' + ret[2][0]
        },
        subtitle: {
          text: "Current Conc is " + ret[1][0]
        },
        xAxis: {
          categories: t
        },
        yAxis: {
          title: {
            text: 'Air Index'
          }
        },
        series: [{
          name: 'PM2.5',
          color: '#FF0000',
          data: p,
          type: 'line'
        }, {
          name: 'Conc',
          color: '#000000',
          data: c,
          type: 'column'
        }]
      });
    });
  });
};
