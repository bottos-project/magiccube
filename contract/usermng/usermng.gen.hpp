#pragma once
#include <eoslib/types.hpp>
#include <eoslib/datastream.hpp>
#include <eoslib/raw_fwd.hpp>

namespace eosio { namespace raw {
   
    template<typename Stream> inline void pack( Stream& s, const user_basic_Info& value ) {
        raw::pack(s, value.info);
    }
    template<typename Stream> inline void unpack( Stream& s, user_basic_Info& value ) {
        raw::unpack(s, value.info);
    }
    template<typename Stream> inline void pack( Stream& s, const reg_user_req& value ) {
        raw::pack(s, value.user_name);
        raw::pack(s, value.info);
    }
    template<typename Stream> inline void unpack( Stream& s, reg_user_req& value ) {
        raw::unpack(s, value.user_name);
        raw::unpack(s, value.info);
    }

    template<typename Stream> inline void pack( Stream& s, const user_login& value ) {
        raw::pack(s, value.user_name);
        raw::pack(s, value.random_num);
    }
    template<typename Stream> inline void unpack( Stream& s, user_login& value ) {
        raw::unpack(s, value.user_name);
        raw::unpack(s, value.random_num);
    }
} }

#include <eoslib/raw.hpp>
namespace eosio {
    void print_ident(int n){while(n-->0){print("  ");}};
    template<typename Type>
    Type current_message_ex() {
        uint32_t size = message_size();
        char* data = (char *)eosio::malloc(size);
        assert(data && read_message(data, size) == size, "error reading message");
        Type value;
        eosio::raw::unpack(data, size, value);
        eosio::free(data);
        return value;
    }

    void dump(const user_basic_Info& value, int tab=0) {
        print_ident(tab);print("info:[");prints(value.info.get_data());print("]\n");      
    }
    template<>
    user_basic_Info current_message<user_basic_Info>() {
        return current_message_ex<user_basic_Info>();
    }
    void dump(const reg_user_req& value, int tab=0) {
        print_ident(tab);print("user_name:[");prints(value.user_name.get_data());print("]\n");
        print_ident(tab);print("info:[");print("\n"); eosio::dump(value.info, tab+1);print_ident(tab);print("]\n");
    }
    template<>
    reg_user_req current_message<reg_user_req>() {
        return current_message_ex<reg_user_req>();
    }
    
    template<>
    user_login current_message<user_login>() {
        return current_message_ex<user_login>();
    }
} //eosio

