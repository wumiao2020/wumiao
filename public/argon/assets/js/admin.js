$(function () {
    $('#modal-default').on('show.bs.modal', function (event) {
        var button = $(event.relatedTarget); // 触发事件的按钮
        var url = button.data('url'); // 解析出data-whatever内容
        var title = button.data('title'); // 解析出data-whatever内容
        var modal = $(this);

        $.ajax({
            cache: true,
            type: "get",
            url: url,
            //data:$('#yourformid').serialize(),// 你的formid
            async: false,
            error: function (request) {
                setTimeout(md.notification('Connection error', 'rose'), 500);
            },
            success: function (data) {
                console.log(data)
                if (data.code === 401){
                    window.location.reload()
                }
                modal.find('.card-header h3').text(title);
                modal.find('.card-body').html(data);
            }
        });
    });

    $('#modal-save').on('click', function (event) {
        var modal = $(this);
        var url = $('#modal-default form').attr('action');
        $.ajax({
            cache: true,
            type: $('#modal-default form').attr('method'),
            url: url,
            data: $('#modal-default form').serialize(),
            async: false,
            dataType: 'json',
            error: function (data) {
                if (data.status !== 'true') {

                } else {
                    setTimeout(md.notification('Connection error', 'rose'), 500);
                }
            },
            success: function (data) {
                modal.find('.modal-body').html(data);
                if (data.status == true) {
                    setTimeout(md.notification(data.message), 500);
                    $('#modal-default').modal('hide')
                } else {
                    $.each(data.responseJSON.errors, function (i, mes) {
                        $.each(mes, function (v, item) {
                            if (item) {
                                setTimeout(md.notification(item, 'rose'), (v + 1) * 500);
                            }
                        })
                    });
                }

            }
        });
    });

});
