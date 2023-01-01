PROTOCOL_PLUGIN_DIRS=$(wildcard ./protocol/*) #网络协议解析插件目录
NOTICE_PLUGIN_DIRS=$(wildcard ./notice/*) #通知服务插件目录

BUILD_SYS=local #编译方式：  local，本地环境编译；linux，linux交叉编译

all: build-plugins

clean: clean-plugins

build-plugins: $(PROTOCOL_PLUGIN_DIRS) $(NOTICE_PLUGIN_DIRS)

clean-plugins: 
	rm -f ./built/*

$(PROTOCOL_PLUGIN_DIRS):
	$(info Clubber plugins at: $(PROTOCOL_PLUGIN_DIRS))
	$(MAKE) $(BUILD_SYS) -C $@

$(NOTICE_PLUGIN_DIRS):
	$(info Clubber plugins at: $(NOTICE_PLUGIN_DIRS))
	$(MAKE) $(BUILD_SYS) -C $@

.PHONY: all $(PROTOCOL_PLUGIN_DIRS) $(NOTICE_PLUGIN_DIRS)