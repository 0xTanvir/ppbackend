/**
 * Created by Tanvir on 2017-07-25.
 */

google.charts.load('current', {'packages':['corechart']});
google.charts.setOnLoadCallback(drawChart);

function drawChart() {
    var data = google.visualization.arrayToDataTable([
        ['Date', 'Bazaar'],
        ['01',  100 ],
        ['02',  110],
        ['03',  66],
        ['04',  100],
        ['05',  200],
        ['06',  250],
        ['07',  100],
        ['08',  250],
        ['09',  250],
        ['10',  250],
        ['11',  150],
        ['12',  870],
        ['13',  150],
        ['14',  150],
        ['15',  150],
        ['16',  654],
        ['17',  150],
        ['18',  150],
        ['19',  759],
        ['20',  150],
        ['21',  150],
        ['22',  150],
        ['23',  800],
        ['24',  150],
        ['25',  150],
        ['26',  150],
        ['27',  187],
        ['28',  150],
        ['29',  150],
        ['30',  499],
        ['31',  150]
    ]);

    var options = {
        title: 'Bazaar Cost',
        hAxis: {title: 'July',  titleTextStyle: {color: '#333'}},
        vAxis: {minValue: 0}
    };

    var chart = new google.visualization.AreaChart(document.getElementById('chart_div'));
    chart.draw(data, options);
}
$(window).resize(function () {
    drawChart();
});
