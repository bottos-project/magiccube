#pragma once
#include <eoslib/types.hpp>
#include <eoslib/datastream.hpp>
#include <eoslib/raw_fwd.hpp>

namespace eosio { namespace raw {
   
    template<typename Stream> inline void pack( Stream& s, const data_file_info& value ) {
        raw::pack(s, value.user_name  );
        raw::pack(s, value.session_id );
        raw::pack(s, value.file_size  );
        raw::pack(s, value.file_name  );
        raw::pack(s, value.file_policy);
        raw::pack(s, value.file_number);
        raw::pack(s, value.signature  );
    }
    template<typename Stream> inline void unpack( Stream& s, data_file_info& value ) {
        raw::unpack(s, value.user_name  );
        raw::unpack(s, value.session_id );
        raw::unpack(s, value.file_size  );
        raw::unpack(s, value.file_name  );
        raw::unpack(s, value.file_policy);
        raw::unpack(s, value.file_number);
        raw::unpack(s, value.signature  );
    }
    template<typename Stream> inline void pack( Stream& s, const reg_data_file_req& value ) {
        raw::pack(s, value.file_hash);
        raw::pack(s, value.info);
    }
    template<typename Stream> inline void unpack( Stream& s, reg_data_file_req& value ) {
        raw::unpack(s, value.file_hash);
        raw::unpack(s, value.info);
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

    void dump(const data_file_info& value, int tab=0) {
        print_ident(tab);print("user_name:  [");prints(value.user_name.get_data());print("]\n");      
        print_ident(tab);print("session_id: [");prints(value.session_id.get_data());print("]\n");   
        print_ident(tab);print("file_size:  [");print(value.file_size);print("]\n");   
        print_ident(tab);print("file_name:  [");prints(value.file_name.get_data());print("]\n");   
        print_ident(tab);print("file_policy:[");prints(value.file_policy.get_data());print("]\n");   
        print_ident(tab);print("file_number:[");print(value.file_number);print("]\n");   
        print_ident(tab);print("signature:  [");prints(value.signature.get_data());print("]\n");   
    }
    template<>
    data_file_info current_message<data_file_info>() {
        return current_message_ex<data_file_info>();
    }
    void dump(const reg_data_file_req& value, int tab=0) {
        print_ident(tab);print("file_hash:[");prints(value.file_hash.get_data());print("]\n");
        print_ident(tab);print("info:     [");print("\n"); eosio::dump(value.info, tab+1);print_ident(tab);print("]\n");
    }
    template<>
    reg_data_file_req current_message<reg_data_file_req>() {
        return current_message_ex<reg_data_file_req>();
    }
    
} //eosio