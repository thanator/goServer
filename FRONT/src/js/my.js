$(document).ready(function () {
    $("#first_button").click(function () {
        $.ajax({
            type: 'GET',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/hi", success: function (result) {
                alert(result)
            }
        });
    });

    $("#go_to_order").click(function () {
        $.ajax({
            type: 'GET',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/hi", success: function (result) {
                alert(result)
            }
        });
    });

    $("#accept_manager").click(function () {
        $.ajax({
            type: 'GET',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/hi", success: function (result) {
                alert(result)
            }
        });
    });

    $("#deny_manager").click(function () {
        $.ajax({
            type: 'GET',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/hi", success: function (result) {
                alert(result)
            }
        });
    });

    $("#see_all_stock_boss").click(function () {
        $.ajax({
            type: 'GET',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/hi", success: function (result) {
                $("#text_area_boss").html(result + " 1")
            }
        });
    });

    $("#see_all_archive_boss").click(function () {
        $.ajax({
            type: 'GET',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/hi", success: function (result) {
                $("#text_area_boss").html(result + " 2")
            }
        });
    });

    $("#manager_req").click(function () {
        $.ajax({
            type: 'GET',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/hi", success: function (result) {
                alert(result)
            }
        });
    });
});

