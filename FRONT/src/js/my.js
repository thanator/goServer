$(document).ready(function () {
    $("#first_button").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/hi", success: function (result) {
                alert(result)
            }
        });
    });

    $("#accept_manager").click(function () {
        var id = $("#dropdown").find("option:selected").text()

        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/accept_manager" + "&" + id, success: function (result) {
                alert(result)
            }
        });
    });

    $("#deny_manager").click(function () {
        var id = $("#dropdown").find("option:selected").text()

        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/deny_manager" + "&" + id , success: function (result) {
                alert(result)
            }
        });
    });

    $("#see_all_stock_boss").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/hi", success: function (result) {
                $("#text_area_boss").html(result + " 1")

            }
        });
    });

    $("#see_all_archive_boss").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/see_all_archive_boss", success: function (result) {
                $("#text_area_boss").html(result)
            }
        });
    });

    $("#manager_req").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/manager_req", success: function (result) {
                $("#dropdown").find('option').remove();
                var arr = result.split(",")
                for (i = 0; i < arr.length; i++) {
                    $("#dropdown").append('<option value="' + arr[i] + '">' + arr[i] + '</option>>')
                }
            }
        });
    });

    $("#manager_find").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/manager_find", success: function (result) {
                $("#manager_text").html(result)
            }
        });
    });
});

