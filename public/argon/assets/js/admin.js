(function ($) {
    $.fn.serializeJson = function () {
        var serializeObj = {};
        var array = this.serializeArray();
        $(array).each(function () {

            if (this.value === 'true') {
                this.value = true;
            } else if (this.value === 'false') {
                this.value = false;
            } else if (!isNaN(this.value)) {
                this.value = parseInt(this.value);
            }

            if (serializeObj[this.name]) {
                if ($.isArray(serializeObj[this.name])) {
                    serializeObj[this.name].push(this.value);
                } else {
                    serializeObj[this.name] = [serializeObj[this.name], this.value];
                }
            } else {
                serializeObj[this.name] = this.value;
            }
        });
        return serializeObj;
    };
})(jQuery);

walter = {
    notification: function (message = "", type = 'success', align = 'right', from = 'top', icon = "ni ni-bell-55") {
        $.notify({
            icon: icon,
            // title: title,
            message: message,
            url: ''
        }, {
            element: 'body',
            type: type,
            allow_dismiss: true,
            placement: {
                from: from,
                align: align
            },
            offset: {
                x: 15, // Keep this as default
                y: 15 // Unless there'll be alignment issues as this value is targeted in CSS
            },
            spacing: 10,
            z_index: 999999,
            delay: 2500,
            timer: 3e3,
            url_target: '_blank',
            mouse_over: false,
            animate: {
                enter: 'animated fadeInRight',
                exit: 'animated fadeOutRight'
            },
            template: '<div data-notify="container" class="alert alert-dismissible alert-{0} alert-notify" role="alert">' +
                '<span class="alert-icon" data-notify="icon"></span> ' +
                '<div class="alert-text"</div> ' +
                '<span class="alert-title" data-notify="title">{1}</span> ' +
                '<span data-notify="message">{2}</span>' +
                '</div>' +
                '<button type="button" class="close" data-notify="dismiss" aria-label="Close"><span aria-hidden="true">&times;</span></button>' +
                '</div>'
        });
    },
    //判断是否为对象（仅为对象，不是数组也不是null）
    isObject: function (exp) {
        return Object.prototype.toString.call(exp) == '[object Object]'
    },

//判断是否为数组（仅为数组，不是对象也不是null）
    isArray: function (exp) {
        return Object.prototype.toString.call(exp) == '[object Array]'
    },

//判断是否为字符串
    isString: function (exp) {
        return Object.prototype.toString.call(exp) == '[object String]'
    },

//判断是否为数字（包括整数和实数）
    isNumber: function (exp) {
        return Object.prototype.toString.call(exp) == '[object Number]'
    },

//判断是否为null
    isNull: function (exp) {
        return Object.prototype.toString.call(exp) == '[object Null]'
    },

//判断是否为undefined
    isUndefined: function (exp) {
        return Object.prototype.toString.call(exp) == '[object Undefined]'
    }
};
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
                setTimeout(walter.notification('Connection error', 'rose'), 500);
            },
            success: function (data) {
                if (data.code === 401) {
                    setTimeout(walter.notification('Login time out', 'rose'), 500);
                    setTimeout(window.location.reload(), 1500);
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
            data: JSON.stringify($('#modal-default form').serializeJson()),
            async: false,
            dataType: 'json',
            contentType: 'application/json;charset=UTF-8',
            error: function (data) {
                if (data.status !== 'true') {
                    walter.notification(data.message,'rose')
                } else {
                    walter.notification('Connection error', 'rose')
                }
            },
            success: function (data) {
                modal.find('.modal-body').html(data);
                if (data.status === true) {
                    walter.notification(data.message)
                    $('#datatable').DataTable().ajax.reload( null, false );
                    $('#modal-default').modal('hide')
                } else {
                    if (walter.isObject(data.message)){
                        walter.notification(data.message.Message,'rose')
                    }
                    walter.notification(data.message,'rose')
                }
            }
        });
    });
    $('#datatable').on('xhr.dt', function ( e, settings, json, xhr ) {
            if (json.code === 401) {
                walter.notification(json.message, 'rose')
                setTimeout(function () {
                    window.location.reload()
                }, 3500);
            }
        } )
});
