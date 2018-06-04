$(document).ready(function () {
    $("#first_button").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/hi",
            success: function (result) {
                alert(result);
            }
        });
    });

    $("#accept_manager").click(function () {
        var id = $("#dropdown").find("option:selected").text();

        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/accept_manager" + "&" + id,
            success: function (result) {
                alert(result);
            }
        });
    });

    $("#deny_manager").click(function () {
        var id = $("#dropdown").find("option:selected").text();

        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/deny_manager" + "&" + id,
            success: function (result) {
                alert(result);
            }
        });
    });

    $("#see_all_stock_boss").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/see_all_stock_boss",
            success: function (result) {
                $("#text_area_boss").html(result);

            }
        });
    });

    $("#see_all_archive_boss").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/see_all_archive_boss",
            success: function (result) {
                $("#text_area_boss").html(result);
            }
        });
    });

    $("#manager_req").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/manager_req",
            success: function (result) {
                $("#dropdown").find('option').remove();
                var arr = result.split(",");
                for (i = 0; i < arr.length; i++) {
                    $("#dropdown").append('<option value="' + arr[i] + '">' + arr[i] + '</option>>');
                }
            }
        });
    });


    $("#boss_find_all_products_id").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/boss_find_all_products_id",
            success: function (result) {
                $("#boss_dropdown_prod").find('option').remove();
                var arr = result.split(",");
                for (i = 0; i < arr.length; i++) {
                    $("#boss_dropdown_prod").append('<option value="' + arr[i] + '">' + arr[i] + '</option>>');
                }
            }
        });
    });


    $("#boss_dropdown_prod").change(function () {
        var value = $(this).find("option:selected").attr("value");
        if (value.split("-").length === 2) {
            $("#boss_find_prod").css('z-index', '-10');
            $("#boss_delete").css('z-index', '10');
        } else {
            $("#boss_find_prod").css('z-index', '10');
            $("#boss_delete").css('z-index', '-10');
        }
    });


    $("#boss_find_all_archive_id").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/boss_find_all_archive_id",
            success: function (result) {
                $("#boss_dropdown_archive").find('option').remove();
                var arr = result.split(",");
                for (i = 0; i < arr.length; i++) {
                    $("#boss_dropdown_archive").append('<option value="' + arr[i] + '">' + arr[i] + '</option>>');
                }
            }
        });
    });


    $("#manager_find").click(function () {

        var id = $("#dropdown").find("option:selected").text();
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/manager_find" + "&" + id,
            success: function (result) {
                $("#manager_text").html(result);
            }
        });
    });

    $("#boss_find_order").click(function () {

        var id = $("#boss_dropdown_archive").find("option:selected").text();
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/boss_find_order" + "&" + id,
            success: function (result) {

                var status = result.split("_");

                $("#text_area_boss").html(status[1]);
                switch (status[0]) {
                    case "1": // waiting
                        //document.getElementById('#good_butt').hidden = true;
                        $('#good_butt').removeClass('hidden');
                        $('#bad_butt').removeClass('hidden');
                        break;
                    case "2": // accepted           
                        $('#good_butt').addClass("hidden");
                        $('#bad_butt').addClass("hidden");
                        break;
                    case "3": // declined          
                        $('#good_butt').removeClass('hidden');
                        $('#bad_butt').addClass("hidden");
                        break;
                    default:
                        break;
                }

            }
        });
    });

    $("#boss_find_prod").click(function () {

        var id = $("#boss_dropdown_prod").find("option:selected").text();
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/boss_find_prod" + "&" + id,
            success: function (result) {
                $("#text_area_boss").html(result);
            }
        });
    });

    $("#boss_delete").click(function () {

        var id = $("#boss_dropdown_prod").find("option:selected").text();
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/boss_delete" + "&" + id,
            success: function (result) {
                //alert(result);
            }
        });
    });

    // region visitor pattern

    $("#boss_export_xml").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/xmlForBoss",
            success: function (result) {
                alert("Экспорт в 'datbos' произведён успешно!");
            }
        });
    });

    $("#manager_export_xml").click(function () {
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/xmlForManager",
            success: function (result) {
                alert("Экспорт в 'datmanag' произведён успешно!");
            }
        });
    });

    //enregion visitor pattern


    // region state pattern

    $("#bad_butt").click(function () {
        var id = $("#boss_dropdown_archive").find("option:selected").text();
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/boss_decline" + "&" + id,
            success: function (result) {
                // alert(result);
            }
        });
    });

    $("#good_butt").click(function () {
        var id = $("#boss_dropdown_archive").find("option:selected").text();
        $.ajax({
            type: 'POST',
            crossdomain: true,
            dataType: 'text',
            url: "http://localhost:3000/boss_accept" + "&" + id,
                    success: function (result) {
                //  alert(result);
            }
        });
    });
    
    //enregion state pattern
    
});