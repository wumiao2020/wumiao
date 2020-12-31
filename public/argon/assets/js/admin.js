(function ($) {
    'use strict';

    const Notify = (function () {
        function notify(placement, align, icon, type, animIn, animOut) {
            $.notify({
                icon: "ni ni-bell-55",
                title: ' Bootstrap Notify',
                message: 'Turning standard Bootstrap alerts into awesome notifications',
                url: ''
            }, {
                element: 'body',
                type: type,
                allow_dismiss: true,
                placement: {
                    from: "top",
                    align: "right"
                },
                offset: {
                    x: 15, // Keep this as default
                    y: 15 // Unless there'll be alignment issues as this value is targeted in CSS
                },
                spacing: 10,
                z_index: 1080,
                delay: 2500,
                timer: 5000,
                url_target: '_blank',
                mouse_over: false,
                animate: {
                    // enter: animIn,
                    // exit: animOut
                    enter: animIn,
                    exit: animOut
                },
                template: '<div data-notify="container" class="alert alert-dismissible alert-{0} alert-notify" role="alert">' +
                    '<span class="alert-icon" data-notify="icon"></span> ' +
                    '<div class="alert-text"</div> ' +
                    '<span class="alert-title" data-notify="title">{1}</span> ' +
                    '<span data-notify="message">{2}</span>' +
                    '</div>' +
                    // '<div class="progress" data-notify="progressbar">' +
                    // '<div class="progress-bar progress-bar-{0}" role="progressbar" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100" style="width: 0%;"></div>' +
                    // '</div>' +
                    // '<a href="{3}" target="{4}" data-notify="url"></a>' +
                    '<button type="button" class="close" data-notify="dismiss" aria-label="Close"><span aria-hidden="true">&times;</span></button>' +
                    '</div>'
            });
        }

    })();

    $(function () {
        $('#modal').on('show.bs.modal', function (event) {
            var button = $(event.relatedTarget) // 触发事件的按钮
            var url = button.data('url') // 解析出data-whatever内容
            var type = button.data('type') // 解析出data-whatever内容
            var title = button.data('title') // 解析出data-whatever内容
            var modal = $(this);

            if (type == 'destroy') {
                $('#modal-save').hide();
                $('#modal-destroy').show();
            } else {
                $('#modal-destroy').hide();
                $('#modal-save').show();
            }

            $.ajax({
                cache: true,
                type: "get",
                url: url,
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                //data:$('#yourformid').serialize(),// 你的formid
                async: false,
                error: function (request) {
                    setTimeout(md.notification('Connection error', 'rose'), 500);
                },
                success: function (data) {
                    modal.find('.modal-title').text(title);
                    modal.find('.modal-body').html(data);
                }
            });
        });

        $('#modal-save').on('click', function (event) {
            var modal = $(this);
            var url = $('#modal form').attr('action');
            $.ajax({
                cache: true,
                type: $('#modal form').attr('method'),
                url: url,
                data: $('#modal form').serialize(),
                async: false,
                dataType: 'json',
                error: function (data) {
                    if (data.status !== 'true') {
                        $.each(data.responseJSON.errors, function (i, mes) {
                            $.each(mes, function (v, item) {
                                if (item) {
                                    setTimeout(md.notification(i + ' ： ' + item, 'rose'), (v + 1) * 500);
                                }
                            })
                        });
                    } else {
                        setTimeout(md.notification('Connection error', 'rose'), 500);
                    }
                },
                success: function (data) {
                    modal.find('.modal-body').html(data);
                    if (data.status == true) {
                        setTimeout(md.notification(data.message), 500);
                        $('#modal').modal('hide')
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

        $('#modalDel').on('show.bs.modal', function (event) {
            var button = $(event.relatedTarget);
            var url = button.data('url');
            var title = button.data('title');
            $(this).find('.modal-title').text(title);
            $('#modalDel-destroy').attr('data-url', url);
        });

        $('#modalDel-destroy').on('click', function (event) {

            alert($('#modalDel-destroy').data('url'));
            console.log($(event.relatedTarget).data('url'))
            console.log($(event.relatedTarget))
            $.ajax({
                cache: true,
                type: 'DELETE',
                url: $('#modalDel-destroy').data('url'),
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                dataType: 'json',
                error: function (data) {
                    $('#modalDel').modal('hide')
                    if (data.responseJSON.message) {
                        setTimeout(md.notification(data.responseJSON.message, 'rose'), 500);
                    }
                    if (data.status !== 'true') {
                        $.each(data.responseJSON.errors, function (i, mes) {
                            $.each(mes, function (v, item) {
                                if (item) {
                                    setTimeout(md.notification(i + ' ： ' + item, 'rose'), (v + 1) * 500);
                                }
                            })
                        });
                    } else {
                        setTimeout(md.notification('Connection error', 'rose'), 500);
                    }
                },
                success: function (data) {
                    $('#modalDel').modal('hide')
                    if (data.status == true) {
                        setTimeout(md.notification(data.message), 500);
                    } else {
                        setTimeout(md.notification(data.message, 'rose'), 500);
                    }

                }
            });
        });

        $('.click').on('click', function () {
            var url = $(this).data('url');
            $.post(url, function (data) {
                if (data.status == 'success') {
                    alert(data.mes);
                    $(this).remove();
                }
            }, "json");
        });
    });
});
