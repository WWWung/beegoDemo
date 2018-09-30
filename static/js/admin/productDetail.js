var laydate
layui.use('laydate', function() {
    laydate = layui.laydate;
    laydate.render({
        elem: "#createTime",
        type: "date"
    })
    laydate.render({
        elem: "#updateTime",
        type: "date"
    })

    tinymce.init({
        selector: "#textarea",
        language: "zh_CN",
        plugins: 'link lists image table colorpicker textcolor wordcount contextmenu codesample',
        fontsize_formats: '10px 11px 12px 14px 16px 18px 20px 24px',
        toolbar: 'bold italic underline strikethrough | fontsizeselect | forecolor backcolor | alignleft aligncenter alignright alignjustify | bullist numlist | outdent indent blockquote | undo redo | link unlink image | removeformat ',
        images_upload_handler: function(blobInfo, success, failure) {
            if (blobInfo.blob().size > 3148576) {
                failure('文件体积过大')
            }
            var formData = new FormData()
            formData.append('file', blobInfo.blob(), blobInfo.filename())
            $.ajax({
                url: "/upload?type=imgs",
                type: 'POST',
                cache: false,
                data: formData,
                processData: false,
                contentType: false,
                beforeSend: function() {
                    uploading = true;
                },
                success: function(d) {
                    success(d.data.src)
                }
            })
        }
    })
})